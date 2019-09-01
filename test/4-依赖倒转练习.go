package main

import "fmt"

type Card interface {
	Display()
}

type Memory interface {
	Storage()
}

type Cpu interface {
	Calculate()
}


type IntelCard struct {
}

func (intel *IntelCard) Display()  {
	fmt.Println("Intel Card display...")
}

type IntelMem struct {
}
func (intel *IntelMem) Storage()  {
	fmt.Println("Intel Memory storage...")
}

type IntelCpu struct {
}
func (intel *IntelCpu) Calculate()  {
	fmt.Println("Intel Cpu calculate...")
}

type NVIDIACard struct {
}

func (nv *NVIDIACard) Display()  {
	fmt.Println("NVIDIA Card display...")
}

type KingstonMem struct {
}
func (ks *KingstonMem) Storage()  {
	fmt.Println("Kingston Memory storage...")
}

//type Compuse interface {
//	Run(card Card, mem Memory, cpu Cpu)
//}

type Computer struct {
	card Card
	memory Memory
	cpu Cpu
}

func NewComputer(card Card, mem Memory, cpu Cpu) *Computer {
	return &Computer{
		card:   card,
		memory: mem,
		cpu:    cpu,
	}
}

func (cp *Computer)Run()  {
	cp.card.Display()
	cp.memory.Storage()
	cp.cpu.Calculate()
}

func main() {
	cp1 := NewComputer(&IntelCard{}, &IntelMem{}, &IntelCpu{})
	cp1.Run()

	cp2 := NewComputer(&NVIDIACard{}, &KingstonMem{}, &IntelCpu{})
	cp2.Run()
}
