package router

import (
	"github.com/gin-gonic/gin"
	wsApi "ku-chat/internal/app/ws"
	"ku-chat/internal/consts"
	"ku-chat/internal/websocket"
)

const WsPAth = "ws"

func RegisterWsRouter(engine *gin.Engine) {
	core := websocket.Core
	core.MsgHandler.AddRouter(consts.FirstTimeOnline, wsApi.Online)
	engine.GET(WsPAth, func(c *gin.Context) { core.Handler(c) })
}
