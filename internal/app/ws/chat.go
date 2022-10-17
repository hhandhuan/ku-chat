package ws

import (
	"ku-chat/internal/websocket"
)

// Online 上线通知
func Online(r *websocket.Request) {
	count := len(r.GetConnection().Core.Connections)
	for _, conn := range r.GetConnection().Core.Connections {
		_ = conn.Send(websocket.Data{ID: 100, Data: count})
	}
}

// Send 发送消息
func Send(r *websocket.Request) {

}
