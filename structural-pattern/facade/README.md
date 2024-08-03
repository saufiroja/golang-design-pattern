# Facade Pattern

Facade pattern adalah sebuah structural pattern yang menyediakan sebuah interface yang simple untuk sebuah body of code yang besar. Hal ini digunakan untuk membuat sebuah system kompleks menadi lebih accessible, mudah digunakan dan lebih dimengerti untuk end user.

## Facade pattern work

Facade pattern menyediakan sebuah interface untuk sekumpulan inteface dalam subsystem. Ini mendefinisikan sebuah high-level interface yang membuat subsystem lebih mudah digunakan.
Subsystem dapat berupa sekumpulan class yang kompleks atau library yang melakukan tugas tertentu. Facade pattern berguna untuk hide complexity dari subsystem dan menyediakan sebuah interface yang lebih simple untuk end user.

## Code Example

```go
package main

import "fmt"

type CPU struct {}

func (c *CPU) Freeze() {
    fmt.Println("Freezing CPU")
}

func (c *CPU) Jump(position string) {
    fmt.Printf("CPU Jumping to: %s\n", position)
}

func (c *CPU) Execute() {
    fmt.Println("CPU Executing")
}

type Memory struct {}

func (m *Memory) Load(position string, data string) {
    fmt.Printf("Memory loading data: %s to position: %s\n", data, position)
}

type HardDrive struct {}

func (h *HardDrive) Read(position string, size int) string {
    data := fmt.Sprintf("HardDrive reading data from position: %s with size: %d\n", position, size)
    fmt.Println(data)

    return data
}

type ComputerFacade struct {
    cpu *CPU
    memory *Memory
    hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        cpu: &CPU{},
        memory: &Memory{},
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
```

Pada contoh diatas, kita memiliki sebuah `ComputerFacade` yang menyediakan sebuah interface simple untuk mengumpulkan subsystem seperti `CPU`, `Memory` dan `HardDrive`. `Start` method pada `ComputerFacade` melakukan serangaikan tidakan yang diperlukan untuk memulai sebuah komputer. End user hanya membutuhkan call `Start` method pada `ComputerFacade` untuk memulai komputer, dan `ComputerFacade` akan menangani kompleksitas dari subsystem yang ada.
