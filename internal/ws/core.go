package ws

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	Core = newCore()
	id   uint32
	up   = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type core struct {
	Connects   map[uint32]*Connection
	connLock   sync.RWMutex
	MsgHandler *Handler
}

func newCore() *core {
	return &core{
		Connects:   make(map[uint32]*Connection),
		MsgHandler: NewHandler(),
	}
}

// Add 新增某个链接
func (c *core) Add(conn *Connection) {
	c.connLock.Lock()
	c.Connects[conn.ID] = conn
	c.connLock.Unlock()
}

// Remove 删除某个链接
func (c *core) Remove(conn *Connection) {
	c.connLock.Lock()
	delete(c.Connects, conn.ID)
	c.connLock.Unlock()
}

// Get 获取某个链接
func (c *core) Get(connID uint32) (*Connection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.Connects[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

// Handler websocket handler
func (c *core) Handler(w http.ResponseWriter, r *http.Request, responseHeader http.Header) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	id++
	connection := NewConnection(id, conn, c)
	c.Add(connection)
	connection.Start()
}
