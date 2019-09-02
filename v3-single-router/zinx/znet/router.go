package znet

import (
	"zinxProject/v3-single-router/zinx/iface"
)

type Router struct {

}

func (r *Router) PreHandle(req iface.IRequest) {
	//fmt.Println("Router PreHandle...")
}
func (r *Router) Handle(req iface.IRequest) {
	//fmt.Println("Router  Handle...")
}
func (r *Router) PostHandle(req iface.IRequest) {
	//fmt.Println("Router PostHandle...")
}