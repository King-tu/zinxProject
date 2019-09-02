package iface

import (
	"net"
)

type IConnect interface {
	Start()
	Stop()
	Send([]byte) (int, error)
	GetTCPConn() *net.TCPConn
	GetConnId() uint32
}

type HandleFunc func(IRequest)

