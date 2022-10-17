package websocket

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
	Connects   map[string]*Connection
	connLock   sync.RWMutex
	MsgHandler *Handler
}

func newCore() *core {
	return &core{
		Connects:   make(map[string]*Connection),
		MsgHandler: NewHandler(),
	}
}

// Add add a conn
func (c *core) Add(conn *Connection) {
	c.connLock.Lock()
	c.Connects[conn.CID] = conn
	c.connLock.Unlock()
}

// Remove remove a conn
func (c *core) Remove(conn *Connection) {
	c.connLock.Lock()
	delete(c.Connects, conn.CID)
	c.connLock.Unlock()
}

// Get get a conn
func (c *core) Get(connID string) (*Connection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.Connects[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// Handler websocket handler
func (c *core) Handler(ctx *gin.Context) {
	conn, err := up.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	connection := NewConn(ctx.Query("cid"), conn, c)
	c.Add(connection)
	connection.Start()
}
