// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package complier

import (
	"web_complier/internal/app/biz"
	"web_complier/internal/app/store"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type ComplierController struct {
	b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore) *ComplierController {
	return &ComplierController{b: biz.NewBiz(ds)}
}
