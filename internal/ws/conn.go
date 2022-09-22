package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Connection struct {
	ID       uint32
	core     *core
	conn     *websocket.Conn
	sendChan chan []byte
	handler  *Handler
}

func NewConnection(ID uint32, Conn *websocket.Conn, core *core) *Connection {
	return &Connection{
		ID:       ID,
		conn:     Conn,
		core:     core,
		sendChan: make(chan []byte),
		handler:  core.MsgHandler,
	}
}

func (c *Connection) reader() {
	defer func() {
		close(c.sendChan)
		c.core.Remove(c)
	}()
	for {
		if _, msg, err := c.conn.ReadMessage(); err != nil {
			if c.isUnexpectedCloseError(err) {
				log.Printf("error: %v", err)
			}
			break
		} else {
			c.sendChan <- msg
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
		case b, ok := <-c.sendChan:
			if !ok {
				break
			}
			var msg Msg
			if err := json.Unmarshal(b, &msg); err != nil {
				log.Println(err)
				continue
			}
			req := &Request{Msg: &msg, Conn: c}
			c.handler.Do(req)
		case <-ticker.C:
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				break
			}
		}
	}
}

func (c *Connection) Start() {
	go c.reader()
	go c.writer()
}
