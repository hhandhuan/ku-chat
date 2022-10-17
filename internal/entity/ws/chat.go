package ws

import "ku-chat/internal/websocket"

type user struct {
	Cid      string `json:"cid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type GroupMsg struct {
	websocket.MsgID
	User user `json:"user"`
}

type OnlineMsg struct {
	ID    int `json:"id"`
	Count int `json:"count"`
}
