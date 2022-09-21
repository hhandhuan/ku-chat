package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Connection struct {
	core     *Core
	ID       uint32
	conn     *websocket.Conn
	sendChan chan []byte
}

func NewConnection(ID uint32, Conn *websocket.Conn, core *Core) *Connection {
	return &Connection{
		core:     core,
		ID:       ID,
		conn:     Conn,
		sendChan: make(chan []byte),
	}
}

func (c *Connection) Reader() {
	defer func() {
		close(c.sendChan)
	}()
	for {
		if msgType, msg, err := c.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		} else {
			fmt.Printf("recv client message: %s type: %d", msg, msgType)
			c.sendChan <- []byte("hello world")
		}
	}
}

func (c *Connection) Writer() {
	for {
		select {
		case msg, ok := <-c.sendChan:
			if !ok {
				break
			}
			for _, connection := range c.core.connects {
				err := connection.conn.WriteMessage(1, msg)
				if err != nil {
					fmt.Println("client write message error")
				}
			}
		}
	}
}
