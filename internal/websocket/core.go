package websocket

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"ku-chat/internal/consts"
	"log"
	"net/http"
	"sync"
)

var (
	Core = newCore()
	up   = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type core struct {
	Disconnects chan *Connection
	Connections map[string]*Connection
	connLock    sync.RWMutex
	MsgHandler  *Handler
}

func newCore() *core {
	return &core{
		Disconnects: make(chan *Connection),
		Connections: make(map[string]*Connection),
		MsgHandler:  NewHandler(),
	}
}

// Add add a conn
func (c *core) Add(conn *Connection) {
	log.Println("add")
	c.connLock.Lock()
	c.Connections[conn.CID] = conn
	c.connLock.Unlock()
	log.Println("end")
}

// Remove remove a conn
func (c *core) Remove(conn *Connection) {
	c.connLock.Lock()
	delete(c.Connections, conn.CID)
	c.connLock.Unlock()
}

// Get get a conn
func (c *core) Get(connID string) (*Connection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.Connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// Quit conn quit
func (c *core) Quit() {
	for {
		select {
		case exitConn := <-c.Disconnects:
			c.Remove(exitConn)
			count := len(c.Connections)
			for _, conn := range c.Connections {
				_ = conn.Send(Data{ID: consts.UserOfflineMsgID, Data: count})
			}
		}
	}
}

// Handler websocket handler
func (c *core) Handler(ctx *gin.Context) {
	conn, err := up.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	go c.Quit()

	connection := NewConn(ctx.Query("cid"), conn, c)
	c.Add(connection)

	go connection.Start()
}
