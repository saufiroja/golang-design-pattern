# Singleton Pattern

Singleton Pattern adalah sebuah design pattern yang membatasi instansiasi sebuah tipe pada satu object. Hal ini berguna ketika hanya satu object yang dibutuhkan untuk mengkoordinasikan tugas di seluruh sistem.

## Singleton Pattern pada Go

Singleton Pattern pada Go dapat diimplementasikan menggunakan sebuah struct dan variable tingkat private yang menampung instance dari struct tersebut. Fungsi constructor struct diekspor dan digunakan untuk mengembalikan instance dari struct tersebut. Berikut adalah contoh implementasi Singleton Pattern pada Go:

```go
package singleton

type singleton struct {
    name string
}

var instance *singleton

func GetInstance() *singleton {
    if instance == nil {
        instance = &singleton{
            name: "Singleton",
        }
    }
    return instance
}

func (s *singleton) GetName() string {
    return s.name
}
```

Pada implementasi di atas, kita telah mendefinisikan sebuah struct `singleton` dengan sebuah field `name`, Kita juga declare package level private bernama `instance` dari tipe `*singleton`. Fungsi `GetInstance` ngereturn instance dari struct `singleton`. Jika `instance` bernilai `nil`, maka fungsi ini akan membuat `instance` baru dari structure `singleton`, fungsi ini akan membuat `instance` baru dari struct `singleton` dan menugaskannya ke variable `instance`. Fungsi `GetName` digunakan untuk mengembalikan nilai dari field `name` pada struct `singleton`.

## Example usage

```go
package main

func main() {
	instance1 := GetInstance("instance1")
	instance2 := GetInstance("instance2")

	if instance1 == instance2 {
		println("Same instance")
	} else {
		println("Different instance")
	}

	println(instance1.GetName())
	println(instance2.GetName())
}
```

## Thread-safe Singleton Pattern

Untuk membuat singleton pattern menjadi thread-safe, kita bisa menggunakan `sync.Lock`. Berikut adalah contoh implementasi thread-safe singleton pattern pada Go:

```go
package singleton

import "sync"

type singleton struct {
    name string
}

var instance *singleton
var lock = &sync.Mutex{}

func GetInstance() *singleton {
    lock.Lock()
    defer lock.Unlock()

    if instance == nil {
        instance = &singleton{
            name: "Singleton",
        }
    }
    return instance
}

func (s *singleton) GetName() string {
    return s.name
}
```

## Use Cases

- Database Connection
  Singleton pattern dapat digunakan untuk membuat koneksi ke database yang bersifat global dan dapat diakses oleh seluruh sistem.
- Logger
  Singleton pattern dapat digunakan untuk membuat logger yang bersifat global dan dapat diakses oleh seluruh sistem. Hal ini berguna untuk menghindari pembuatan logger yang berlebihan.
