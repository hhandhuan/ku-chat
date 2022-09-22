package ws

import "ku-chat/internal/ws"

type UserAuthMsg struct {
	ws.MsgID
	Data struct {
		UUID     string `json:"uuid"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	} `json:"data"`
}
