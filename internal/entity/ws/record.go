package ws

import "ku-chat/internal/model"

type AddFriendWsReq struct {
	ID   uint32 `json:"id"`
	Data struct {
		User   model.Users `json:"user"`
		Remark string      `json:"remark"`
	} `json:"data"`
}

type RejectFriendWsReq struct {
	ID   uint32 `json:"id"`
	Data struct {
		User   model.Users `json:"user"`
		Remark string      `json:"remark"`
	} `json:"data"`
}
