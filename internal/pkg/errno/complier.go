// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package errno

var (
	// ErrLangNotFound 不支持该语言
	ErrLangNotFound = &Errno{HTTP: 404, Code: "FailedRun.ErrLangNotFound", Message: "lang was not found"}
)
