# Bridge Pattern

Bridge pattern adalah sebuah structural design pattern yang memisahkan abstraksi dari implementasinya sehingga keduanya dapat bervariasi secara independen. Pattern ini
berguna ketika kita memiliki abstraksi yang dapat memiliki beberapa implementasi, dan kita ingin dapat beralih di antara mereka pada saat runtime.
Bridge pattern menyediakan cara untuk mengatur code sehingga detail implementasi disembunyikan dari klien.

## Code Example

Misalkan kita memiliki interface `Shape` dan 2 implementasi dari interface tersebut yaitu `Circle` dan `Square`. Selain itu, misalkan kita memiliki 2 inteface `Renderer` yang
mengimplementasikan rendering bentuk-bentuk ini: `VectorRenderer` dan `RasterRenderer`. Untuk mengimplementasikan bridge pattern, kita membuat sebuah bridge antara interface
`Shape` dan `Renderer`. Hal ini memungkinkan kita untuk beralih di antara berbagai `renderer` tanpa mengubah interface `Shape`.

```go
type Renderer interface {
    RenderCircle(radius int)
    RenderSquare(side int)
}

type Shape interface {
    Draw()
}

type Circle struct {
    x, y, radius float64
    renderer Renderer
}

func (c *Circle) Draw() {
    c.renderer.RenderCircle(c.radius)
}

type Square struct {
    side int
    renderer Renderer
}

func (s *Square) Draw() {
    s.renderer.RenderSquare(s.side)
}

type VectorRenderer struct {}

func (v *VectorRenderer) RenderCircle(radius int) {
    fmt.Println("Drawing a circle of radius", radius)
}

func (v *VectorRenderer) RenderSquare(side int) {
    fmt.Println("Drawing a square of side", side)
}

type RasterRenderer struct {}

func (r *RasterRenderer) RenderCircle(radius int) {
    fmt.Println("Drawing pixels for a circle of radius", radius)
}

func (r *RasterRenderer) RenderSquare(side int) {
    fmt.Println("Drawing pixels for a square of side", side)
}
```

Dalam conttoh di atas, kita memiliki inteface `Shape` dan 2 implementasi dari interface tersebut yaitu: `Circle` dan `Square`.
Baik `Circle` dan `Square` memiliki sebuah renderer yang mengimplementasikan interface `Renderer`.
`Renderer` ini memiliki 2 metode yaitu `RenderCircle` dan `RenderSquare`. Kita juga memiliki 2 implementasi dari interface `Renderer` yaitu `VectorRenderer` dan `RasterRenderer`.

```go
package main

import "fmt"

func main() {
    vectorRenderer := &VectorRenderer{}
    circle := &Circle{0, 0, 5, vectorRenderer}
    square := &Square{5, vectorRenderer}

    circle.Draw()
    square.Draw()

    rasterRenderer := &RasterRenderer{}
    circle.renderer = rasterRenderer
    square.renderer = rasterRenderer

    circle.Draw()
    square.Draw()
}
```

## Kelebihan dan Kekurangan Bridge Pattern

1. Seperaion of concerns: Bridge pattern memisahkan antara abstraksi dari implementasinya, yang membuatnya lebih mudah memodifikasi dan memelihara code.
2. Encapsulation: Bridge pattern mengEncapsulation detail implementasi, yang membuat code lebih aman dan tidak terlalu rentan terhadap kesalahan.
3. Flexibility: Bridge pattern memungkinkan pembuatan implementasi baru tanpa mempengaruhi abstraksi.
4. Reusability: Bridge pattern mendorong penggunaan ulang code dengan mengizinkan abstraksi yang berbeda untuk menggunakan implementasi yang sama.
5. Testability: Bridge Pattern membuat kode lebih mudah diuji dengan mengizinkan implementasi untuk ditiru.

## Perbedaan Bridge Pattern dengan Adapter Pattern

### Adapter Pattern

Ketika kita ingin mengadaptasi interface yang sudah ada untuk memenuhi kebutukan client. Kita melakukan ini dengan membuat interface baru yang dapat digunakan oleh client,
dan mengimplementasikan interface ini dengan mengadaptasi interface yang sudah ada. Dengan kata lain, adapter pattern mengubah interface yang sudah ada agar dapat digunakan oleh client.

### Bridge Pattern

Ketika kita ingin memisahkan antar abstraksi dari impelmentasinya sehingga keduanya dapat bervariasi secara independen. Kita melakukan ini dengan membuat bridge antara abstraksi dan implementasi. Dengan kata lain, bridge pattern menciptakan lapisan abstraksi baru antara client dan implementasi.
