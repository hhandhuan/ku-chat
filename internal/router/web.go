package router

import (
	"github.com/gin-gonic/gin"
	"ku-chat/internal/app/web"
)

func RegisterWebRouter(engine *gin.Engine) {
	engine.GET("/login", web.Login)
	engine.POST("/login", web.Login)

	engine.GET("/register", web.Register)
	engine.POST("/register", web.Register)

	engine.Use(auth)

	engine.GET("/", web.Home)
	engine.GET("logout", web.Logout)
}
