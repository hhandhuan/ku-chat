package ws

import (
	"ku-chat/internal/ws"
	"log"
)

func Ping(r *ws.Request) {
	log.Println(r.GetData())
}

func Test(r *ws.Request) {
	log.Println(r.GetData())
}
