package web

type AddReq struct {
	TargetID uint32 `v:"required#用户名错误" form:"target_id"`
	Remark   string `form:"remark"`
}
