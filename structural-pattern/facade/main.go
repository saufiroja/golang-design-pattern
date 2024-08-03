package main

import "fmt"

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("Freezing CPU")
}

func (c *CPU) Jump(position string) {
	fmt.Printf("CPU Jumping to: %s\n", position)
}

func (c *CPU) Execute() {
	fmt.Println("CPU Executing")
}

type Memory struct{}

func (m *Memory) Load(position string, data string) {
	fmt.Printf("Memory loading data: %s to position: %s\n", data, position)
}

type HardDrive struct{}

func (h *HardDrive) Read(position string, size int) string {
	data := fmt.Sprintf("HardDrive reading data from position: %s with size: %d\n", position, size)
	fmt.Println(data)

	return data
}

type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	c.cpu.Freeze()
	c.memory.Load("0x00", c.hardDrive.Read("0x00", 1024))
	c.cpu.Jump("0x00")
	c.cpu.Execute()
}

func main() {
	computer := NewComputerFacade()
	computer.Start()
}
