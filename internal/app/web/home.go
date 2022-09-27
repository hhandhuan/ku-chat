package web

import (
	"github.com/gin-gonic/gin"
	"ku-chat/internal/model"
	"ku-chat/internal/service"
	"log"
)

func Home(ctx *gin.Context) {
	s := service.Context(ctx)

	user := s.Auth()

	var unread int64
	// 获取用户未读消息
	res := model.Record().M.Where("target_id", user.ID).Where("state", 0).Count(&unread)
	if res.Error != nil {
		log.Println(res.Error)
	}

	s.View("home", gin.H{"unread": unread})
}
