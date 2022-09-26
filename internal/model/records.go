package model

import (
	"gorm.io/gorm"
	"ku-chat/pkg/db"
	"time"
)

type Records struct {
	Model
	UserId   int64      `gorm:"column:user_id" json:"user_id"`     //用户ID
	TargetId int64      `gorm:"column:target_id" json:"target_id"` //目标用户ID
	State    int8       `gorm:"column:state" json:"state"`         //状态: 0-申请中/1-已处理/2-已拒绝
	Remark   string     `gorm:"column:remark" json:"remark"`       //申请备注
	ReadedAt *time.Time `gorm:"column:readed_at" json:"readed_at"` //阅读时间
}

type RecordModel struct {
	M     *gorm.DB
	Table string
}

func Record() *RecordModel {
	return &RecordModel{M: db.DB.Model(&Records{}), Table: "records"}
}
