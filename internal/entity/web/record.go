package web

import "ku-chat/internal/model"

type AddReq struct {
	TargetID uint32 `v:"required#用户名错误" form:"target_id"`
	Remark   string `form:"remark"`
}

type RecordLog struct {
	model.Records
	User model.Users `gorm:"foreignKey:user_id"`
}
