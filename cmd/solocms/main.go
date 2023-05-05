// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/solocms

package main

import (
	"os"
	"solocms/internal/solocms"

	_ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := solocms.NewSoloCMSCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}