package model

import (
	"gorm.io/gorm"
	"ku-chat/pkg/db"
)

type Friends struct {
	Model
	UserId   uint64 `gorm:"column:user_id" json:"user_id"`     //用户 ID
	TargetId uint64 `gorm:"column:target_id" json:"target_id"` //目标 ID
}

type FriendModel struct {
	M     *gorm.DB
	Table string
}

func Friend() *FriendModel {
	return &FriendModel{M: db.DB.Model(&Friends{}), Table: "friends"}
}
