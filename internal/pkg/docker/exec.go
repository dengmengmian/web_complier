// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package docker

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	"web_complier/internal/pkg/log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func DockerRun(image string, code string, dest string, cmd string, langTimeout int64, memory int64, ext string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Errorw("NewClientWithOpts:", err)
	}

	optionFilters := filters.NewArgs()

	optionFilters.Add("name", image)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: optionFilters,
		Limit:   1,
	})
	if err != nil {
		log.Errorw("docker list err:", err)
	}
	var containerID string
	if len(containers) > 0 {
		filtersContainer := containers[0]
		containerID = containers[0].ID
		if filtersContainer.State == "exited" {
			if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
				log.Errorw("ContainerStart err:%v:", err)
			}
		}
	} else {
		bindsString := fmt.Sprintf("%s:%s", "/tmp", dest)
		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image:        image,
			AttachStderr: true,
			AttachStdout: true,
			Tty:          true,
		}, &container.HostConfig{
			Binds: []string{bindsString},
			Resources: container.Resources{
				Memory: memory, // Minimum memory limit allowed is 6MB.
			},
		}, nil, nil, fmt.Sprintf("%s", image)) //并发创建容器会报错，但可以保证只有一个容器

		if err != nil {
			log.Errorw("ContainerCreate:", err)
		}
		containerID = resp.ID
		if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
			log.Errorw("ContainerStart err:%v:", err)
		}
	}

	rand.Seed(time.Now().UnixMicro())
	filename := fmt.Sprintf("test_%d", rand.Uint32())
	fname := fmt.Sprintf("%s/%s.%s", "/tmp", filename, ext)

	err = os.WriteFile(fname, []byte(code), 0777)
	if err != nil {
		log.Errorw("write file err:", err)
	}

	cmd = strings.ReplaceAll(cmd, "filename", dest+"/"+filename)
	cmd = fmt.Sprintf("timeout %d %s > %s/%s.log", langTimeout, cmd, dest, filename)
	res, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd: []string{"sh", "-c", cmd},
	})
	if err != nil {
		removeFile(filename)
		log.Errorw("docker exec create err:", err)
	}

	chanC := make(chan int64)
	var resSting string
	go func() {
		t1 := time.Now().UnixMicro()
		if err := cli.ContainerExecStart(ctx, res.ID, types.ExecStartCheck{Detach: false, Tty: false}); err != nil {
			removeFile(filename)
			log.Errorw("ContainerExecStart %d err:%v:", res.ID, err)
		}

		logFile := fmt.Sprintf("%s/%s.log", "/tmp", filename)
		for tryTimes := langTimeout * 100; tryTimes > 0; tryTimes-- {
			time.Sleep(time.Duration(20) * time.Millisecond)
			dir, err := os.Stat(logFile)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				log.Errorw("open log file err:", err)
			}
			if dir.Size() > 0 {
				content, err := os.ReadFile(logFile)
				if err != nil {
					log.Errorw("read log file err", err)
				}
				resSting = string(content)
				break
			}
		}
		timeCost := time.Now().UnixMicro() - t1
		chanC <- timeCost
	}()

	timeout := time.NewTimer(time.Duration(langTimeout) * time.Second)
	timeoutFlag := false
	select {
	case <-chanC:
		break
	case <-timeout.C:
		timeoutFlag = true
		log.Errorw("execute timeout")
		return "execute timeout"
	}
	removeFile(filename)
	if timeoutFlag {
		resSting = resSting + "\n execute timeout"
	}
	return resSting
}

func removeFile(filename string) {
	pattern := fmt.Sprintf("%s/%s*", "/tmp", filename)
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Infow("removeFile", err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Errorw("removeFile", err)

		}
	}
}
