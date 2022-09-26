package service

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ku-chat/pkg/config"
	"net/http"

	"ku-chat/internal/model"
)

const (
	errKey     = "err"
	msgKey     = "msg"
	dataKey    = "data"
	userKey    = "user"
	versionKey = "version"
)

func Context(ctx *gin.Context) *BaseContext {
	return &BaseContext{
		Ctx:     ctx,
		session: sessions.Default(ctx),
		path:    "/",
		Conf:    config.Conf,
	}
}

type BaseContext struct {
	Ctx     *gin.Context
	session sessions.Session
	path    string
	Conf    *config.Config
}

// Redirect 处理跳转
func (c *BaseContext) Redirect() {
	c.Ctx.Redirect(http.StatusFound, c.path)
}

func (c *BaseContext) clear() {
	c.session.Delete(errKey)
	c.session.Delete(msgKey)
	_ = c.session.Save()
}

// Back 返回上一页
func (c *BaseContext) Back() *BaseContext {
	c.path = c.Ctx.Request.RequestURI
	return c
}

// To 设置跳转路径
func (c *BaseContext) To(to string) *BaseContext {
	c.path = to
	return c
}

// WithError 错误信息跳转
func (c *BaseContext) WithError(err interface{}) *BaseContext {
	errStr := ""
	switch v := err.(type) {
	case error:
		errStr = v.Error()
	case string:
		errStr = v
	}
	c.session.Set(errKey, errStr)
	_ = c.session.Save()
	return c
}

// WithMsg 提示消息跳转
func (c *BaseContext) WithMsg(msg string) *BaseContext {
	c.session.Set(msgKey, msg)
	_ = c.session.Save()
	return c
}

// SetAuth 设置授权
func (c *BaseContext) SetAuth(users model.Users) {
	s, _ := json.Marshal(users)
	c.session.Set(userKey, string(s))
	_ = c.session.Save()
}

// Auth 获取授权
func (c *BaseContext) Auth() *model.Users {
	var user *model.Users
	str := c.session.Get(userKey)
	if str == nil {
		return user
	}
	if v, ok := str.(string); ok {
		_ = json.Unmarshal([]byte(v), &user)
	}
	return user
}

// Check 检查授权
func (c *BaseContext) Check() bool {
	user := c.Auth()
	if user == nil {
		return false
	} else {
		return user.ID > 0
	}
}

// Forget 清除授权
func (c *BaseContext) Forget() {
	c.session.Delete(userKey)
	_ = c.session.Save()
}

// View 模版返回
func (c *BaseContext) View(tpl string, data interface{}) {
	obj := gin.H{
		versionKey: config.Conf.App.Version,
		errKey:     c.session.Get(errKey),
		msgKey:     c.session.Get(msgKey),
		userKey:    c.Auth(),
		dataKey:    data,
	}
	c.clear()
	c.Ctx.HTML(http.StatusOK, tpl, obj)
}

// Json 通用 JSON 响应
func (c *BaseContext) Json(data interface{}) {
	c.Ctx.JSON(http.StatusOK, data)
}
