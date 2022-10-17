package router

import (
	"github.com/gin-gonic/gin"
	wsApi "ku-chat/internal/app/ws"
	"ku-chat/internal/consts"
	"ku-chat/internal/websocket"
)

func RegisterWsRouter(engine *gin.Engine) {
	core := websocket.Core

	core.MsgHandler.AddRouter(consts.UserOnlineMSgID, wsApi.Online)
	core.MsgHandler.AddRouter(consts.SendGroupMsgID, wsApi.Send)
	
	engine.GET("ws", func(c *gin.Context) { core.Handler(c) })
}
