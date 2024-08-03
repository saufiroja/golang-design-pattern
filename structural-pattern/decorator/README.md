# Decorator Pattern

Decorator Pattern adalah sebuah design pattern yang memungkinkan perilaku ditambahkan ke sebuah object, secara dinamis, tanpa mempengaruhi perilaku dari object lain dari class yang sama.

## Example

Dalam Go, decorator pattern bisa diimplementasikan dengan menggunakan interface dan anonymous function. Berikut adalah contoh implementasi dari decorator pattern.

```go
package main

import "fmt"

type Printer interface {
    Print() string
}

type SimplePrinter struct {
}

func (sp *SimplePrinter) Print() string {
    return "Simple Printer"
}

func BoldDecorator(p Printer) Printer {
    return PrinterFunc(func() string {
        return fmt.Sprintf("<b>%s</b>", p.Print())
    })
}

type PrinterFunc func() string

func (pf PrinterFunc) Print() string {
    return pf()
}

func main() {
    simplePrinter := &SimplePrinter{}
    boldPrinter := BoldDecorator(simplePrinter)

    fmt.Println(simplePrinter.Print()) // Simple Printer
    fmt.Println(boldPrinter.Print())  // <b>Simple Printer</b>
}
```

## Middleware Decorator

Another example dimana kita bisa menggunakan decorator pattern adalah pada middleware di web server. Middleware function adalah function yang execute sebelum atau sesudah sebuah request di handle oleh sebuah web server. Mereka bisa berguna untuk task seperti logging, authentication, dll.

Dalam go, kita bisa menggunakan decorator patterm untuk membuat middleware function. Kita bisa membuuat sebuah `Handler` interface yang representasikan sebuah function yang handler sebuah HTTP request dan return sebuah HTTP Status code dan sebuah response body. Kemudian kita bisa membuat sebuah `Middleware` inteface yang menerima sebuah `Handler` dan return `Handler` lainnya yang executes beberapa perilaku tambahan sebelum atau sesudah handler utama di execute.

```go
package main

import (
    "fmt"
    "net/http"
)

type Handler func (r *http.Request) (int, string)

type Middleware func (h Handler) Handler

func LoggingMiddleware(h Handler) Handler {
    return func(r *http.Request) (int, string) {
        fmt.Println("Handling request")
        status, body := h(r)
        fmt.Printf("Request handled with status %d\n", status)
        return status, body
    }
}

func main() {
    handler := func(r *http.Request) (int, string) {
        return http.StatusOK, "Hello, World!"
    }

    logginHandler := LoggingMiddleware(handler)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        status, body := logginHandler(r)
        w.WriteHeader(status)
        w.Write([]byte(body))
    })

    http.ListenAndServe(":8080", nil)
}
```

Dalam contoh diatas, kita membuuat sebuah `Handler` function yang return sebuah HTTP status code dan sebuah response body. Kemudian kita mendefinisikan sebuah `LogginMiddleware` function yang menerima sebuah `Handler` dan return sebuah `Handler` lainnya yang menambahkan log sebelum dan sesudah handler utama di execute.
