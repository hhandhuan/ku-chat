package ws

import (
	ws2 "ku-chat/internal/entity/ws"
	"ku-chat/internal/ws"
	"log"
)

func Auth(r *ws.Request) {
	var msg *ws2.UserAuthMsg
	if err := r.Parse(&msg); err != nil || msg == nil {
		log.Println("ws.api.user.auth: json unmarshal error:", err)
		return
	}
	r.GetConnection().Conn.WriteMessage(1, r.Data)
}
