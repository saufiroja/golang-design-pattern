# Composite Pattern

Composite pattern adalah structural design pattern yang memungkinkan kamu untuk menyusun obeject kedalam sebuah tree structure dan bekerja dengan tree tersebut seolah-olah itu adalah sebuah object tunggal.

Composite pattern berguna ketika kamu membutuhkan untuk bekerja dengan structure hirarki. Itu memungkinkan kamu untuk membuat sebuah struktur seperti tree yang dimana setiap node dalam tree bisa jadi object tunggal atau sebuah group dari object. Semua node dalam tree bisa diperlakukan sebagai object tunggal atau sebagai dari group yang lebih besar.

Dalam composite pattern, terdapat 2 type object: `leaf` dan `composite` object.

- `Leaf` Object adalah last node dalam tree dan tidak memiliki child.
- `Composite` object dapat memiliki 1 atau lebih child node, yang dapat berupa `leaf` object atau `composite` object lainnya.

## Implementasi Composite Pattern dalam Go

Dalam go, kita dapat mengimplementasikan composite pattern dengan membuat sebuah interface `Component` yang akan diimplementasikan oleh semua object dalam tree

```go
type Component interface {
    Operation()
}
```

`Operation` adalah method yang akan diimplementasikan oleh semua object dalam tree. Implementasi dari method ini akan berbeda untuk `leaf` dan `composite` object.

Untuk implementasi `composite` object, kita membuat sebuah struct `Composite` yang memiliki field `children` yang merupakan slice dari `Component`

```go
type Composite struct {
    children []Component
}
```

`Composite` struct akan memiliki list dari child node, yang akan berupa `leaf` atau `composite` object lainnya.

```go
func (c *Composite) Operation() {
    fmt.Println("Composite operation")
    for _, child := range c.children {
        child.Operation()
    }
}
```

`Operation` method dari `Composite` struct akan melakukan operasi pada semua child node.
Untuk implementasi `leaf` object, kita membuat sebuah struct `Leaf` yang menyimpan data untuk node tersebut.

```go
type Leaf struct {
    data string
}
```

`Leaf` struct akan menyimpan data untuk leaf node. Kita kemudian dapat mengimplementasikan method `Operation` untuk `Leaf` struct.

```go
func (l *Leaf) Operation() {
    fmt.Println("Leaf operation with data:", l.data)
}
```

`Operation` method dari `Leaf` struct akan melakukan operasi pada leaf node.

## Contoh Code

```go
func main() {
    root := &Composite{
        children: []Component{
            &Leaf{data: "file1.txt"},
            &Composite{
                children: []Component{
                    &Leaf{data: "file2.txt"},
                    &Composite{
                        children: []Component{
                            &Leaf{data: "file3.txt"},
                        },
                    },
                },
            },
        },
    }

    root.Operation()
}
```

Output:

```
Composite operation
Leaf operation with data: file1.txt
Composite operation
Leaf operation with data: file2.txt
Composite operation
Leaf operation with data: file3.txt
```
