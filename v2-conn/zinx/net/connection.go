package net

import (
	"fmt"
	"log"
	"net"
	"zinxProject/v2-conn/zinx/iface"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	IsClosed bool
	HandleAPI HandleFunc
}

type HandleFunc func(iface.IConnect, []byte)


func NewConnection(conn *net.TCPConn, connID uint32, handleFunc HandleFunc) *Connection {

	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		IsClosed:  false,
		HandleAPI: handleFunc,
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

		handleFunc(c, buf[:cnt])
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
