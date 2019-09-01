package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func init()  {
	log.SetPrefix("Error Msg: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	conn, err := net.Dial("tcp4", ":8848")
	if err != nil {
		log.Fatalln("net.Dial err: ", err)
	}

	data := []byte("hello world")

	for {
		cnt, err := conn.Write(data)
		if err != nil {
			log.Fatalln("net.Write err: ", err)
		}
		fmt.Println("Client ===> Server, len:", cnt, "data:", string(data))

		buf := make([]byte, 1024)
		//read
		cnt, err = conn.Read(buf)
		if err != nil {
			log.Fatalln("net.Read err: ", err)
		}
		fmt.Println("Client <== Server, len:", cnt, "data:", string(buf[:cnt]))

		time.Sleep(time.Second * 1)
	}
}
