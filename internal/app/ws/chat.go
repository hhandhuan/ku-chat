package ws

import (
	"encoding/json"
	"ku-chat/internal/websocket"
	"log"
)

// Online 上线通知
func Online(r *websocket.Request) {
	count := len(r.GetConnection().Core.Connections)
	for _, conn := range r.GetConnection().Core.Connections {
		_ = conn.Send(websocket.Data{ID: 100, Data: count})
	}
}

type MsgReq struct {
	ID   uint32 `json:"id"`
	Data struct {
		User struct {
			CID    string `json:"cid"`
			Name   string `json:"name"`
			Avatar string `json:"avatar"`
		} `json:"user"`
		Content string `json:"content"`
	} `json:"data"`
}

// Send 发送消息
func Send(r *websocket.Request) {
	var msg MsgReq
	log.Println(string(r.GetData()))
	err := json.Unmarshal(r.GetData(), &msg)
	if err != nil {
		log.Printf("json decode error: %v", err)
		return
	}
	for _, conn := range r.GetConnection().Core.Connections {
		if conn.CID == r.GetConnection().CID {
			continue
		} else {
			_ = conn.Send(websocket.Data{ID: 300, Data: msg})
		}
	}
}
