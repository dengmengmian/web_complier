// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/solocms

package model

import (
	"time"

	"solocms/pkg/auth"

	"gorm.io/gorm"
)

type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	Email     string    `gorm:"column:email"`          //邮箱
	Nickname  string    `gorm:"column:nickname"`       //昵称
	Password  string    `gorm:"column:password"`       //密码
	Phone     string    `gorm:"column:phone"`          //手机号
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //更新时间
	CreatedAt time.Time `gorm:"column:createdAt"`      //创建时间
	Username  string    `gorm:"column:username"`       //用户名
}

// TableName sets the insert table name for this struct type
func (u *UserM) TableName() string {
	return "user"
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	// Encrypt the user password.
	u.Password, err = auth.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return nil
}
