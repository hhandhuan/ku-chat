package web

import (
	"github.com/gin-gonic/gin"
	"ku-chat/internal/service"
)

func Home(ctx *gin.Context) {
	service.Context(ctx).View("home", nil)
}
