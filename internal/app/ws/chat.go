package ws

import (
	"ku-chat/internal/websocket"
	"log"
)

func Online(r *websocket.Request) {
	log.Println(r.GetData())
	// 给本地链接群发消息
	for _, connection := range r.GetConnection().Core.Connects {
		if err := connection.SendByte(r.GetData()); err != nil {
			log.Println("sendByte error: ", err)
			continue
		}
	}
}
