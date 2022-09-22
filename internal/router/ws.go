package router

import (
	"github.com/gin-gonic/gin"
	wsApi "ku-chat/internal/api/ws"
	"ku-chat/internal/consts"
	"ku-chat/internal/ws"
)

const WsPAth = "ws"

func RegisterWsRouter(engine *gin.Engine) {
	core := ws.Core
	core.MsgHandler.AddRouter(consts.UserAuthMsgID, wsApi.Auth)
	engine.GET(WsPAth, func(c *gin.Context) { core.Handler(c) })
}
