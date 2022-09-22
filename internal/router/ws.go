package router

import (
	"github.com/gin-gonic/gin"
	wsApi "ku-chat/internal/api/ws"
	"ku-chat/internal/ws"
)

func RegisterWsRouter(engine *gin.Engine) {
	core := ws.Core

	core.MsgHandler.AddRouter(1, wsApi.Ping)
	core.MsgHandler.AddRouter(2, wsApi.Test)

	engine.GET("/ws", func(c *gin.Context) {
		core.Handler(c.Writer, c.Request, nil)
	})
}
