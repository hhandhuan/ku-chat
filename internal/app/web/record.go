package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ku-chat/internal/consts"
	ew "ku-chat/internal/entity/web"
	es "ku-chat/internal/entity/ws"
	"ku-chat/internal/model"
	"ku-chat/internal/service"
	"ku-chat/internal/websocket"
	"time"
)

var Record = cRecord{}

type cRecord struct{}

// Add 添加好友申请
func (*cRecord) Add(ctx *gin.Context) {
	s := service.Context(ctx)

	var req ew.AddReq
	if err := ctx.ShouldBind(&req); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
		return
	}
	if err := g.Validator().Data(req).Run(context.Background()); err != nil {
		s.Json(gin.H{"code": 1, "msg": err.Error()})
		return
	}

	user := s.Auth()

	if user.ID == uint64(req.TargetID) {
		s.Json(gin.H{"code": 1, "msg": "无法添加自身为好友"})
		return
	}

	var record *model.Records
	res := model.Record().M.Where("user_id", user.ID).Where("target_id", req.TargetID).Find(&record)
	if res.Error != nil {
		s.Json(gin.H{"code": 1, "msg": res.Error})
		return
	}

	if record.ID > 0 {
		// TODO 阻止客户端消息数+1
		if record.ReadedAt == nil && (time.Now().Unix()-record.UpdatedAt.Unix()) < 600 {
			s.Json(gin.H{"code": 0, "msg": "申请添加好友成功"})
			return
		}
		res = model.Record().M.Where("id", record.ID).Updates(map[string]interface{}{
			"readed_at": nil,
			"state":     0,
			"remark":    req.Remark,
		})
	} else {
		res = model.Record().M.Create(&model.Records{
			UserId:   gconv.Int64(s.Auth().ID),
			TargetId: gconv.Int64(req.TargetID),
			Remark:   req.Remark,
			State:    0,
		})
	}

	if res.Error != nil {
		s.Json(gin.H{"code": 1, "msg": "申请添加好友失败"})
		return
	}

	// 本地服务器获取链接
	conn, err := websocket.Core.Get(gconv.String(req.TargetID))
	if err != nil {
		s.Json(gin.H{"code": 0, "msg": "申请添加好友成功"})
		return
	}

	// 发送申请好友消息
	wsReq := es.AddFriendWsReq{}
	wsReq.ID = consts.AddFriendWsMsgID
	wsReq.Data.User = *user
	wsReq.Data.Remark = req.Remark

	if err := conn.Send(wsReq); err != nil {
		s.Json(gin.H{"code": 1, "msg": "申请添加好友失败"})
	} else {
		s.Json(gin.H{"code": 0, "msg": "申请添加好友成功"})
	}
}

// Logs 好友申请日志
func (*cRecord) Logs(ctx *gin.Context) {
	s := service.Context(ctx)

	var logs []*ew.RecordLog
	res := model.Record().M.Where("target_id", s.Auth().ID).Order("updated_at DESC").Preload("User").Find(&logs)
	if res.Error != nil {
		s.Json(gin.H{"code": 1, "msg": "获取记录失败"})
		return
	}

	//model.Record().M.Where("target_id", s.Auth().ID).Update("readed_at", time.Now())

	s.Json(gin.H{"code": 0, "data": logs})
}
