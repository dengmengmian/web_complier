// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package complier

import (
	"context"
	"web_complier/internal/app/store"
	"web_complier/internal/pkg/docker"
	"web_complier/internal/pkg/errno"
	"web_complier/internal/pkg/log"
	v1 "web_complier/pkg/api/v1"
)

// ComplierBiz 定义了 complier 模块在 biz 层所实现的方法.
type ComplierBiz interface {
	Run(ctx context.Context, r *v1.RunRequest) (*v1.RunResponse, error)
}

// ComplierBiz 接口的实现.
type complierBiz struct {
	ds store.IStore
}

// 确保 complierBiz 实现了 complierBiz 接口.
var _ ComplierBiz = (*complierBiz)(nil)

// New 创建一个实现了 complierBiz 接口的实例.
func New(ds store.IStore) *complierBiz {
	return &complierBiz{ds: ds}
}

// 调用docker运行代码,并返回结果
func (b *complierBiz) Run(ctx context.Context, r *v1.RunRequest) (*v1.RunResponse, error) {
	langexists, err := docker.LangExists(r.Lang)
	if !langexists {
		log.C(ctx).Errorw("ErrLangNotFound", "err", err)
		return nil, errno.ErrLangNotFound
	}
	tpl := docker.Run(r.Lang)
	output := docker.DockerRun(tpl.Image, r.Code, tpl.Dir, tpl.Cmd, tpl.Timeout, tpl.Memory, tpl.Ext)
	return &v1.RunResponse{Output: output}, nil
}
