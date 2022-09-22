package ws

import (
	"ku-chat/internal/ws"
	"log"
)

func Ping(request *ws.Request) {
	log.Println(request.GetData())
}

func Test(request *ws.Request) {
	log.Println(request.GetData())
}
