package main

import (
	"log"
	"v1-basic-server/zinx/net"
)

func init()  {
	log.SetPrefix("Error Msg: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	server := net.NewServer("zinx v1.0")
	server.Serve()
}
