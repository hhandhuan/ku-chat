package model

import (
	"gorm.io/gorm"
	"ku-chat/pkg/db"
)

type Users struct {
	Model
	Name     string `gorm:"column:name" db:"name" json:"name"`       //用户昵称
	Avatar   string `gorm:"column:avatar" db:"avatar" json:"avatar"` //用户头像
	Password string `gorm:"column:password" db:"password" json:"-"`  //用户密码
}

type userModel struct {
	M     *gorm.DB
	Table string
}

func User() *userModel {
	return &userModel{M: db.DB.Model(&Users{}), Table: "users"}
}
