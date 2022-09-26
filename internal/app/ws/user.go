package ws

import (
	ews "ku-chat/internal/entity/ws"
	"ku-chat/internal/websocket"
	"log"
)

func Auth(r *websocket.Request) {
	var msg *ews.UserAuthMsg
	if err := r.Parse(&msg); err != nil || msg == nil {
		log.Println("ws.api.user.auth: json unmarshal error:", err)
		return
	}
	err := r.GetConnection().Conn.WriteMessage(1, r.Data)
	if err != nil {
		log.Println("ws.api.user.auth: write message error: ", err)
	}
}
