package main

import (
	"log"
	"zinxProject/v2-conn/zinx/net"
)

func init()  {
	log.SetPrefix("Error Msg: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	server := net.NewServer("zinx v1.0")
	server.Serve()
}
