package ws

import (
	"errors"
	"sync"
)

type Core struct {
	connects map[uint32]*Connection
	connLock sync.RWMutex
}

func NewCore() *Core {
	return &Core{connects: make(map[uint32]*Connection)}
}

func (c *Core) Add(conn *Connection) {
	c.connLock.Lock()
	c.connects[conn.ID] = conn
	c.connLock.Unlock()
}

func (c *Core) Remove(conn *Connection) {
	c.connLock.Lock()
	delete(c.connects, conn.ID)
	c.connLock.Unlock()
}

func (c *Core) Get(connID uint32) (*Connection, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.connects[connID]; ok {
		return conn, nil
	}
	return nil, errors.New("connection not found")
}
