# Singleton Pattern

Singleton Pattern adalah sebuah design pattern yang membatasi instansiasi sebuah tipe pada satu object. Hal ini berguna ketika hanya satu object yang dibutuhkan untuk mengkoordinasikan tugas di seluruh sistem.

Singleton memastikan bahwa sebuah class hanya memiliki satu instance dan menyediakan titik akses golbal ke instance tersebut. Dengan kata lain, tidak peduli berapa kali class di panggil, hanya ada 1 object yang akan di buat.

# Problem

Skenario yang umum adalah ketika anda memiliki sebuah configuration atau connection database yang perlu diakses secara global oleh seluruh sistem.
Jika setiap class membuat instansi sendiri dari object tersebut, maka ini bisa menyebabkan konsumsi memory yang berlebih,
konflik configuration atau ketidakkonsistenan dalam data.

Bayangkan dalam aplikasi backend, kamu memiliki configuration database yang ingin digunakan di berbagai bagian aplikasi. Jika setiap kali aplikasi memerlukan koneksi database
dan membuat instansi sendiri, ini bisa menyebabkan masalah seperti terlalu banyak koneksi yang terbuka, konflik dalam konfigurasi, atau kesulitan dalam mengelola lifecycle dari koneksi tersebut.

# Solution

Menggunakan singleton pattern dapat memberikan manfaat berikut:

1. **Pengelolaan Resource**: Singleton memastikan bahwa resource seperti koneksi database atau konfigurasi hanya diinisialisasi sekali dan digunakan oleh seluruh sistem.
2. **Efisiensi Memory**: Dengan hanya membuat 1 instance dari object, memory yang digunakan lebih efisien dibandingkan membuat instance yang sama berulang kali.
3. **Konsistensi Data**: Semua bagian dari aplikasi akan menggunakan instance yang sama, sehingga data yang diakses akan konsisten.

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

# Kelebihan dan Kekurangan

1. Kelebihan
   - Kamu dapat memastikan bahwa sebuah class hanya memiliki 1 instance.
   - Kamu mendapatkan jalur access global ke instance tersebut.
   - Singleton Obect diinisialisasi hanya ketika diminta untuk pertama kali.
2. Kekurangan
   - Melanggar prinsip Single Responsibility Principle (SRP), pattern ini memecahkan 2 masalah sekaligus.
   - Singleton pattern dapat menutupi desain yang buruk. Misalnya, ketika komponen program mengetahui terlalu banyak tentang satu sama lain.
   - Pattern ini membutuhkan perlakuan khusus dalam lingkungan multithreaded sehingga beberapa thread tidak akan membuat singleton object beberapa kali.
   - Mungkin akan sulit untuk di unit test karena banyak kerangka kerja test yang mengandalkan inheritance sata menghasilkan object tiruan.
