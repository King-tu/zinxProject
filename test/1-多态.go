package main

import "fmt"

type Call interface {
	Call()
}

type Nokia struct {

}

func (n *Nokia) Call()  {
	fmt.Println("Nokia call ...")
}

type Apple struct {

}

func (n *Apple) Call()  {
	fmt.Println("Apple call ...")
}

func PhoneCall(c Call)  {
	c.Call()
}

func main() {
	PhoneCall(&Nokia{})
	PhoneCall(&Apple{})
}