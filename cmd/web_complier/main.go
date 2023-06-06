// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package main

import (
	"os"
	web_complier "web_complier/internal/app"

	_ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := web_complier.Newweb_complierCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
