// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package web_complier

import (
	"web_complier/internal/app/controller/v1/complier"
	"web_complier/internal/app/controller/v1/user"
	"web_complier/internal/app/store"
	"web_complier/internal/pkg/core"
	"web_complier/internal/pkg/errno"
	"web_complier/internal/pkg/log"
	"web_complier/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

// installRouters 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)
	cm := complier.New(store.S)

	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(middleware.Authn())
		}

		complierV1 := v1.Group("/complier")
		{
			complierV1.POST("", cm.Run)
		}
	}

	return nil
}
