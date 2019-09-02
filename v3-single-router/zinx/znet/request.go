package znet

import "zinxProject/v3-single-router/zinx/iface"

type Request struct {
	IConn iface.IConnect
	Data []byte
	Len uint32
}

func NewRequest(iconn iface.IConnect, data []byte, len uint32) iface.IRequest {
	return &Request{
		IConn: iconn,
		Data: data,
		Len:  len,
	}
}

func (req *Request) GetIConn() iface.IConnect {
	return req.IConn
}
func (req *Request) GetData() []byte {
	return req.Data
}
func (req *Request) GetDataLen() uint32 {
	return req.Len
}

