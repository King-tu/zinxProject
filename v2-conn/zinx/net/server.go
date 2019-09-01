package net

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
	"zinxProject/v2-conn/zinx/iface"
)

func handleFunc(iconn iface.IConnect, data []byte) {
	retInfo := strings.ToUpper(string(data))

	cnt, err := iconn.Send([]byte(retInfo))
	if err != nil {
		log.Println("iconn.Send err: ", err)
		return
	}

	fmt.Printf("Server ==> Client, len: %d, buf: %s\n", cnt, retInfo)
}

type Server struct {
	IP string
	Port int32
	Name string
	IPVersion string
}

func NewServer(name string) iface.IServer {
	return &Server{
		IP:      "0.0.0.0",
		Port:    8848,
		Name:    name,
		IPVersion: "tcp4",
	}
}

func (s *Server) Start()  {
	fmt.Println("Server start...")

	address := fmt.Sprintf("%s:%d", s.IP, s.Port)
	tcpAddress, err := net.ResolveTCPAddr(s.IPVersion, address)
	if err != nil {
		log.Println("net.ResolveTCPAddr err: ", err)
		return
	}

	tcpListener, err := net.ListenTCP(s.IPVersion, tcpAddress)
	if err != nil {
		log.Println("net.ListenTCP err: ", err)
		return
	}

	go func() {
		for {
			//fmt.Println("等待客户端连接：")
			tcpConn, err := tcpListener.AcceptTCP()
			if err != nil {
				log.Println("tcpListener.AcceptTCP err: ", err)
				//return
				continue
			}

			var connID uint32
			connID = 0
			go func() {
			/*	defer tcpConn.Close()
				remoteAddr := tcpConn.RemoteAddr().String()
				fmt.Printf("客户端 %s 已连接...\n", remoteAddr)*/
				connection := NewConnection(tcpConn, connID, handleFunc)
				connID++

				go connection.Start()
			}()
		}
	}()
}

func (s *Server) Stop()  {
	fmt.Println("Server stop...")

}

func (s *Server) Serve()  {

	s.Start()
	fmt.Println("Server serve...")

	for{
		runtime.GC()
	}
}
