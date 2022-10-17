package model

import (
	"gorm.io/gorm"
	"ku-chat/pkg/db"
	"time"
)

type Records struct {
	Model
	UserId   uint64     `gorm:"column:user_id"`                    //用户ID
	TargetId uint64     `gorm:"column:target_id"`                  //目标用户ID
	State    uint8      `gorm:"column:state"`                      //状态: 0-申请中/1-已处理/2-已拒绝
	Remark   string     `gorm:"column:remark"`                     //申请备注
	ReadedAt *time.Time `gorm:"column:readed_at" json:"readed_at"` //阅读时间
}

type RecordModel struct {
	M     *gorm.DB
	Table string
}

func Record() *RecordModel {
	return &RecordModel{M: db.DB.Model(&Records{}), Table: "records"}
}
