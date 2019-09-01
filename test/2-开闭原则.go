package main

import "fmt"

//抽象的银⾏业务员
type AbstractBanker interface{
	DoBusi() //抽象的处理理业务接⼝口
}
//存款的业务员
type SaveBanker struct {
}
func (sb *SaveBanker) DoBusi() {
	fmt.Println("存款的业务员，操作了存款业务...")
}
//转账的业务员
type TransferBanker struct {
}
func (tb *TransferBanker) DoBusi() {
	fmt.Println("转账的业务员，操作了转账业务...")
}
//⽀付的业务员
type PayBanker struct {
	//AbstractBanker
}
func (pb *PayBanker) DoBusi() {
	fmt.Println("⽀付的业务员，操作了⽀付业务...")
}
func main() {
	//存款业务
	sb := &SaveBanker{}
	sb.DoBusi()
	//转账业务
	tb := &TransferBanker{}
	tb.DoBusi()
	//支付业务
	pb := &PayBanker{}
	pb.DoBusi()
}