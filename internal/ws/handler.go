package ws

import (
	"fmt"
	"log"
	"strconv"
)

type MsgHandleApi func(request *Request)

type Handler struct {
	apis map[uint32]MsgHandleApi
}

func NewHandler() *Handler {
	return &Handler{apis: make(map[uint32]MsgHandleApi)}
}

func (m *Handler) Do(req *Request) {
	if api, ok := m.apis[req.GetMsgID()]; !ok {
		log.Println("msg id not found")
	} else {
		api(req)
	}
}

func (m *Handler) AddRouter(msgID uint32, api MsgHandleApi) {
	if _, ok := m.apis[msgID]; ok {
		panic("repeated api , msgID = " + strconv.Itoa(int(msgID)))
	} else {
		m.apis[msgID] = api
		fmt.Println("add api msgID = ", msgID)
	}
}
