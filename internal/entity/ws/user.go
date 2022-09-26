package ws

import "ku-chat/internal/websocket"

type UserAuthMsg struct {
	websocket.MsgID
	Data struct {
		UUID     string `json:"uuid"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	} `json:"data"`
}
