// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package v1

// LoginRequest 指定了 `POST /run` 接口的请求参数.
type RunRequest struct {
	Code  string `form:"code" json:"code" valid:"required,stringlength(6|1024)"`
	Input string `form:"input" json:"input"`
	Lang  string `form:"lang" json:"lang" valid:"required"`
}

type RunResponse struct {
	Output string `json:"output"`
}
