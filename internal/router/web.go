package router

import (
	"github.com/gin-gonic/gin"
	"ku-chat/internal/app/web"
)

func RegisterWebRouter(engine *gin.Engine) {
	engine.GET("/login", web.User.Login)
	engine.POST("/login", web.User.Login)

	engine.GET("/register", web.User.Register)
	engine.POST("/register", web.User.Register)

	engine.Use(auth)

	engine.GET("/", web.Home)
	engine.GET("logout", web.User.Logout)
	engine.GET("search", web.User.Search)

	engine.POST("record-add", web.Record.Add)
	engine.GET("record-logs", web.Record.Logs)
}
