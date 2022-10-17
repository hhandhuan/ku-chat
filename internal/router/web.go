package router

import (
	"github.com/gin-gonic/gin"
	"ku-chat/internal/app/web"
)

func RegisterWebRouter(engine *gin.Engine) {
	engine.Any("/login", web.User.Login)
	engine.Any("/register", web.User.Register)
	engine.Use(auth)
	engine.GET("/", web.Home)
	engine.Any("logout", web.User.Logout)
}
