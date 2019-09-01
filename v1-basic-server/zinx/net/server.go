package net

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strings"
	"zinxProject/v1-basic-server/zinx/iface"
)

type Server struct {
	IP string
	Port int32
	Name string
	Version string
}

func NewServer(name string) iface.IServer {
	return &Server{
		IP:      "0.0.0.0",
		Port:    8848,
		Name:    name,
		Version: "tcp4",
	}
}

func (s *Server) Start()  {
	fmt.Println("Server start...")

	address := fmt.Sprintf("%s:%d", s.IP, s.Port)
	tcpAddress, err := net.ResolveTCPAddr(s.Version, address)
	if err != nil {
		log.Println("net.ResolveTCPAddr err: ", err)
		return
	}

	tcpListener, err := net.ListenTCP(s.Version, tcpAddress)
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
				return
			}
			go func() {
				defer tcpConn.Close()
				remoteAddr := tcpConn.RemoteAddr().String()
				fmt.Printf("客户端 %s 已连接...\n", remoteAddr)
				for {
					buf := make([]byte, 1024)
					cnt, err := tcpConn.Read(buf)
					if err != nil {
						if err == io.EOF {
							fmt.Printf("客户端 %s 已断开连接...\n", remoteAddr)
							break
						} else {
							log.Println("tcpConn.Read err: ", err)
							return
						}
					}
					fmt.Printf("Server <== Client, len: %d, buf: %s\n", cnt, string(buf[:cnt]))

					retInfo := strings.ToUpper(string(buf[:cnt]))
					cnt, err = tcpConn.Write([]byte(retInfo))
					if err != nil {
						log.Println("tcpConn.Write err: ", err)
						return
					}
					fmt.Printf("Server ==> Client, len: %d, buf: %s\n", cnt, retInfo)
				}
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
