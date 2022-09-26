package router

import (
	"github.com/gin-gonic/gin"
	"ku-chat/internal/service"
)

func auth(ctx *gin.Context) {
	s := service.Context(ctx)
	if s.Check() {
		ctx.Next()
	} else {
		s.To("login").WithError("请登录后在操作").Redirect()
		ctx.Abort()
		return
	}
}
