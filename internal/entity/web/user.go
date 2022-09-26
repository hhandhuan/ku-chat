package web

type RegisterReq struct {
	Name            string `v:"required|regex:^[\u4e00-\u9fa5a-zA-Z0-9]{2,8}$#用户名错误|用户名格式错误" form:"name"`
	Password        string `v:"required|length:6,20|same:confirm_password#密码错误|密码长度错误|密码和确认密码不一致" form:"password"`
	ConfirmPassword string `v:"required|length:6,20#确认密码错误|密码长度错误" form:"confirm_password"`
}

type LoginReq struct {
	Name     string `v:"required|regex:^[\u4e00-\u9fa5a-zA-Z0-9]{2,8}$#用户名错误|用户名格式错误" form:"name"`
	Password string `v:"required|length:6,20#密码错误|密码长度错误" form:"password"`
}
