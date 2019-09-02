package main

import (
	"fmt"
	"log"
	"strings"
	"zinxProject/v3-single-router/zinx/iface"
	"zinxProject/v3-single-router/zinx/znet"
)

func init()  {
	log.SetPrefix("Error Msg: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type myRouter struct {
	znet.Router
}

/*func (r *myRouter) PreHandle(request iface.IRequest) {
	fmt.Println("Router PreHandle...")
}*/

func (r *myRouter) Handle(request iface.IRequest) {
	fmt.Println("用户自定义Handle() called...")
	iconn := request.GetIConn()
	data := request.GetData()

	retInfo := strings.ToUpper(string(data))
	retInfo = retInfo + "你好！"

	cnt, err := iconn.Send([]byte(retInfo))
	if err != nil {
		log.Println("iconn.Send err: ", err)
		return
	}

	fmt.Printf("Server ==> Client, len: %d, buf: %s\n", cnt, retInfo)
}
/*func (r *myRouter) PostHandle(request iface.IRequest) {
	fmt.Println("Router PostHandle...")
}
*/




func main() {
	server := znet.NewServer("zinx v1.0")
	server.AddRouter(&myRouter{})
	server.Serve()
}
