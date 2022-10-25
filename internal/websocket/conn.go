package websocket

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Connection struct {
	CID      string
	Core     *core
	Conn     *websocket.Conn
	SendChan chan []byte
	Handler  *Handler
	Ctx      *gin.Context
	ExitChan chan bool
}

func NewConn(CID string, Conn *websocket.Conn, core *core) *Connection {
	return &Connection{
		CID:      CID,
		Conn:     Conn,
		Core:     core,
		SendChan: make(chan []byte),
		Handler:  core.MsgHandler,
		ExitChan: make(chan bool),
	}
}

func (c *Connection) reader() {
	defer func() {
		c.Conn.Close()
		close(c.SendChan)
		close(c.ExitChan)
		log.Println("client close")
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if c.isUnexpectedCloseError(err) {
				log.Printf("error: %v", err)
			}
			c.Core.Disconnects <- c
			c.ExitChan <- true
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
		case bd := <-c.SendChan:
			if id := c.checkMsgID(bd); id <= 0 {
				break
			} else {
				c.Handler.Do(&Request{MsgID: id, Data: bd, Conn: c})
			}
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				break
			}
		case <-c.ExitChan:
			return // select for 中 break 只会跳出 select 不会跳出 for
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

// Send 发送消息给当前链接
func (c *Connection) Send(msg interface{}) error {
	byteData, _ := json.Marshal(msg)
	if err := c.Conn.WriteMessage(1, byteData); err != nil {
		return errors.New("send message error")
	} else {
		return nil
	}
}

// SendByte 发送消息给当前链接
func (c *Connection) SendByte(msg []byte) error {
	if err := c.Conn.WriteMessage(1, msg); err != nil {
		return errors.New("send message error")
	} else {
		return nil
	}
}
