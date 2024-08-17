# Chain Of Responsibility

Chain of Responsibility adalah salah satu behavioral design pattern yang memungkinkan sejumlah object untuk menangani sebuah request tanpa mengetahui obect mana ynag menangani request tersebut. Setiap obejct di dalam chain of resposibility memiliki kesempatan untuk memproses request atau meneruskannya ke object berikutnya dalam chain.

# Problem

Bayangkan kita memiliki system proses request dengan beberapa langkah validasi. Misalnya, kita sedang mengembangkan system management content yang memerlukan validasi input dari user sebelum content di simpan ke dalam database. Validasi ini bisa berupa pengecekan authentication, validasi format data, duplikasi data, dan lain-lain. Tanpa chain of responsibility, mungkin akan banyak if-else statement yang harus kita tulis untuk menangani validasi tersebut.

# Solution

Dengan menggunakan chain of responsibility, kita dapat memisahkan setiap langkah validasi ke dalam komponen yang berbeda. Setiap komponen hanya menangani 1 jenis validasi dan jika validasi berhasil, request akan diteruskan ke komponen berikutnya. Dengan demikian, ini memungkinkan kita untuk menambahkan, menghapus, atau mengubah urutan validasi tanpa mengubah kode yang sudah ada.

# Use Case dalam Go

1. Request Handling Middleware
   Menggunakan pattern ini untuk membangun middleware yang menangani request HTTP seperti authentication, logging, atau validasi input.
2. Error Handling
   Menangani error dalam suatu proses di mana setiap handler di dalam chain mencoba memperbaiki error atau meneruskannya ke handler berikutnya.
3. Command Processing
   Menggunakan chain of responsibility untuk memproses command dengan berbagai handler yang berurutan.

# Contoh Implementasi

```go
type Service interface {
	HelloWorld(name string) (string, error)
}

type service struct{}

func (s *service) HelloWorld(name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

// validator
type validator struct {
	next Service
}

func (v *validator) HelloWorld(name string) (string, error) {
	if len(name) <= 3 {
		return "", fmt.Errorf("Name is too short")
	}

	return v.next.HelloWorld(name)
}

type logger struct {
	next Service
}

func (l *logger) HelloWorld(name string) (string, error) {
	res, err := l.next.HelloWorld(name)
	if err != nil {
		return "", err
	}

	fmt.Println("Request is successful")
	return res, nil
}

func New() Service {
	return &logger{
		next: &validator{
			next: &service{},
		},
	}
}
```

```go
func main() {
	s := New()
	res, err := s.HelloWorld("John")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res, err = s.HelloWorld("Jo")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
```

# Kekurangan dan Kelebihan

1. Kelebihan
   1. Kita dapat mengontrol urutan request handling tanpa mengubah kode yang sudah ada.
   2. Single Responsibility Principle: Kita dapat memisahkan class yang memanggil operasi dari class yang melaksanakan operasi.
   3. Open/Closed Principle: Kita dapat menambahkan atau menghapus handler tanpa mengubah kode yang sudah ada.
2. Kekurangan
   1. Beberapa request mungkin tidak unhandled.
