package znet

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"zinxProject/v3-single-router/zinx/iface"
)

type Server struct {
	IP string
	Port int32
	Name string
	IPVersion string
	Router iface.IRouter
}

/*func handleFunc(request iface.IRequest) {
}*/

func NewServer(name string) iface.IServer {
	return &Server{
		IP:      "0.0.0.0",
		Port:    8848,
		Name:    name,
		IPVersion: "tcp4",
		Router: &Router{},
	}
}

func (s *Server) Start()  {
	fmt.Println("Server start...")

	address := fmt.Sprintf("%s:%d", s.IP, s.Port)
	tcpAddress, err := net.ResolveTCPAddr(s.IPVersion, address)
	if err != nil {
		log.Println("znet.ResolveTCPAddr err: ", err)
		return
	}

	tcpListener, err := net.ListenTCP(s.IPVersion, tcpAddress)
	if err != nil {
		log.Println("znet.ListenTCP err: ", err)
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
				connection := NewConnection(tcpConn, connID, s.Router)
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

func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
}