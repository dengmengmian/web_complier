// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package user

import (
	"web_complier/internal/pkg/core"
	"web_complier/internal/pkg/errno"
	"web_complier/internal/pkg/log"
	v1 "web_complier/pkg/api/v1"

	"github.com/gin-gonic/gin"
)

// 登录 miniblog 并返回一个 JWT Token.
func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
