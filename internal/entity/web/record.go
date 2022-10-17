package web

import "ku-chat/internal/model"

type AddReq struct {
	TargetID uint64 `v:"required" form:"target_id"`
	Remark   string `form:"remark"`
}

type AuditReq struct {
	ID    uint64 `v:"required" form:"id"`
	State uint   `v:"required|in:1,2" form:"state"`
}

type RecordLog struct {
	model.Records
	User model.Users `gorm:"foreignKey:user_id" json:"user"`
}
