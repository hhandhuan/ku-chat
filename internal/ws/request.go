package ws

type Request struct {
	Msg  *Msg
	Conn *Connection
}

func (r *Request) GetConnection() *Connection {
	return r.Conn
}

func (r *Request) GetData() string {
	return r.Msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.Msg.GetID()
}
