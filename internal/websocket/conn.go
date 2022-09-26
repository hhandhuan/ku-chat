package websocket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Connection struct {
	CID      string
	Core     *core
	Conn     *websocket.Conn
	SendChan chan []byte
	Handler  *Handler
	Ctx      *gin.Context
}

func NewConn(CID string, Conn *websocket.Conn, core *core) *Connection {
	return &Connection{
		CID:      CID,
		Conn:     Conn,
		Core:     core,
		SendChan: make(chan []byte),
		Handler:  core.MsgHandler,
	}
}

func (c *Connection) reader() {
	defer func() {
		close(c.SendChan)
		c.Core.Remove(c)
	}()
	log.Println(c.CID)
	for {
		if _, msg, err := c.Conn.ReadMessage(); err != nil {
			if c.isUnexpectedCloseError(err) {
				log.Printf("error: %v", err)
			}
			break
		} else {
			c.SendChan <- msg
		}
	}
}

func (c *Connection) writer() {
	ticker := time.NewTicker(time.Second * 54)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case byteData, ok := <-c.SendChan:
			if !ok {
				break
			}
			msgID := c.checkMsgID(byteData)
			if msgID <= 0 {
				log.Println("msg ID error")
				continue
			}
			c.Handler.Do(&Request{MsgID: msgID, Data: byteData, Conn: c})
		case <-ticker.C:
			err := c.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				break
			}
		}
	}
}

// isUnexpectedCloseError 是意外关闭错误
func (c *Connection) isUnexpectedCloseError(err error) bool {
	return websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway)
}

// Start 开启处理链接
func (c *Connection) Start() {
	go c.reader()
	go c.writer()
}

// checkMsgID 检查消息ID
func (c *Connection) checkMsgID(msg []byte) uint32 {
	var msgID MsgID
	if err := json.Unmarshal(msg, &msgID); err != nil {
		log.Println("json decode error: ", err)
		return 0
	} else {
		return msgID.ID
	}
}
