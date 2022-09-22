package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Connection struct {
	ID       uint32
	Core     *core
	Conn     *websocket.Conn
	SendChan chan []byte
	Handler  *Handler
}

func NewConnection(ID uint32, Conn *websocket.Conn, core *core) *Connection {
	return &Connection{
		ID:       ID,
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

func (c *Connection) isUnexpectedCloseError(err error) bool {
	return websocket.IsUnexpectedCloseError(
		err,
		websocket.CloseAbnormalClosure,
		websocket.CloseGoingAway,
	)
}

func (c *Connection) writer() {
	ticker := time.NewTicker(time.Second * 54)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case b, ok := <-c.SendChan:
			if !ok {
				break
			}
			var msg Msg
			if err := json.Unmarshal(b, &msg); err != nil {
				log.Println(err)
				continue
			}
			req := &Request{Msg: &msg, Conn: c}
			c.Handler.Do(req)
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				break
			}
		}
	}
}

func (c *Connection) Start() {
	go c.reader()
	go c.writer()
}
