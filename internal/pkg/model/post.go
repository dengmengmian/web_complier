// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/web_complier

package model

import "time"

type postM struct {
	ID        int64     `gorm:"column:id;primary_key"` //
	PostID    string    `gorm:"column:postID"`         //文章字符ID
	Content   string    `gorm:"column:content"`        //内容
	Username  string    `gorm:"column:username"`       //作者
	Title     string    `gorm:"column:title"`          //标题
	CreatedAt time.Time `gorm:"column:createdAt"`      //创建时间
	UpdatedAt time.Time `gorm:"column:updatedAt"`      //更新时间

}

// TableName sets the insert table name for this struct type
func (p *postM) TableName() string {
	return "post"
}
