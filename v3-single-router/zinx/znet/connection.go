package znet

import (
	"fmt"
	"log"
	"net"
	"zinxProject/v3-single-router/zinx/iface"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	IsClosed bool
	//HandleAPI iface.HandleFunc
	Router iface.IRouter
}


func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {

	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		IsClosed:  false,
		//HandleAPI: handleFunc,
		Router: router,
	}
}

func (c *Connection) Start() {
	for {
		buf := make([]byte, 1024)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("tcpConn.Read err: ", err)
			return
		}
		fmt.Printf("Server <== Client, len: %d, buf: %s\n", cnt, string(buf[:cnt]))

		req := NewRequest(c, buf[:cnt], uint32(cnt))

		//handleFunc(req)
		c.Router.PreHandle(req)
		c.Router.Handle(req)
		c.Router.PostHandle(req)
	}
}

func (c *Connection) Stop() {
	if !c.IsClosed {
		_ = c.Conn.Close()
	}
}
func (c *Connection) Send(data []byte) (int, error) {
	cnt, err := c.Conn.Write(data)
	if err != nil {
		log.Println("tcpConn.Write err: ", err)
		return -1, err
	}
	return cnt, nil
}

func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}
