package ws

import "ku-chat/internal/websocket"

// GroupMsg 群消息
type GroupMsg struct {
	websocket.MsgID
	Data struct {
		CID      string `json:"cid"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		Content  string `json:"content"`
	} `json:"data"`
}
