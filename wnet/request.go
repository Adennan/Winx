package wnet

import "Winx/wiface"

type Request struct {
	Conn wiface.IConnection
	Data []byte
}

func (r *Request) GetConn() wiface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Data
}
