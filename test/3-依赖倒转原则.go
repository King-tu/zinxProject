package main

import "fmt"

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

type BMW struct {
}

type Benz struct {
}

func (bmw *BMW) Run()  {
	fmt.Println("BMW running .... ")
}

func (benz *Benz) Run()  {
	fmt.Println("Benz running .... ")
}

type Lisi struct {
}

func (lisi *Lisi)Drive(car Car)  {
	fmt.Print("lisi driving..., ")
	car.Run()
}

type Zhangsan struct {
}

func (zs *Zhangsan)Drive(car Car)  {
	fmt.Print("Zhangsan driving..., ")
	car.Run()
}


func main() {
	benz := Benz{}
	bmw := BMW{}

	lisi := Lisi{}
	lisi.Drive(&benz)
	lisi.Drive(&bmw)

	zs := Zhangsan{}
	zs.Drive(&benz)
	zs.Drive(&bmw)
}
