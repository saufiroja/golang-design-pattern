# SOLID Principles

SOLID adalah kumpulan 5 prinsip design yang dimaksudkan untuk membuat design software lebih mudah dipahami, fleksible dan dapat dipelihara.

## S - Single Responsibility Principle

Cobalah untuk membuat setiap class bertanggung jawab untuk 1 bagian fungsionalitas yang disediakan oleh software, dan buatlah tanggung jawab tersebut sepenuhnya dienkapsulasi dalam class tersebut.

Tujuan utama dari prinsip ini adalah mengurangi kompleksitas. Kita tidak perlu menciptakan design yang canggih untuk sebuah program yang hanya memiliki sekita 200 baris code. Buatlah selusin method yang cantik dan mudah dipahami, dan kita akan baik-baik saja.

Jika suatu class melakukan terlalu banyak hal, Kita harus mengubahnya setiap kali salah satu dari hal-hal tersebut berubah. Ketika melakukan hal tersebut, Kita beresiko merusak bagian lain dari class yang bahkan tidak ingin kita ubah.

Jika kita merasa sulit untuk fokus pada aspek tertentu dari program satu per satu, ingatlah prinsip single responsibility dan periksa apakah sudah waktunya untuk membagi beberapa class menjadi beberapa bagian yang lebih kecil.

## Example

Misalkan kita memiliki sebuah aplikasi yang menangani pendaftaran pengguna (user registration). Ada beberapa langkah dalam pendaftaran ini, seperti validasi data pengguna, menyimpan data ke database, dan mengirim email konfirmasi.

- Contoh melanggar prinsip single responsibility

Pada contoh di bawah ini, `UserService` bertanggung jawab untuk melakukan validasi data pengguna, menyimpan data ke database, dan mengirim email konfirmasi. Ini melanggar prinsip single responsibility karena `UserService` melakukan terlalu banyak hal.

```go
package main

import (
	"fmt"
	"net/smtp"
)

type User struct {
	Name  string
	Email string
}

type UserService struct{}

func (s *UserService) Register(user User) error {
	// 1. Validasi data pengguna
	if user.Name == "" || user.Email == "" {
		return fmt.Errorf("invalid user data")
	}

	// 2. Simpan ke database (contoh sederhana)
	fmt.Println("Saving user to database...")

	// 3. Kirim email konfirmasi
	auth := smtp.PlainAuth("", "example@gmail.com", "password", "smtp.example.com")
	to := []string{user.Email}
	msg := []byte("Subject: Registration Successful!\n\nWelcome to our service!")
	err := smtp.SendMail("smtp.example.com:587", auth, "example@gmail.com", to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("User registered successfully!")
	return nil
}

func main() {
	userService := UserService{}
	user := User{Name: "John Doe", Email: "john.doe@example.com"}
	if err := userService.Register(user); err != nil {
		fmt.Println("Error:", err)
	}
}
```

- Contoh menerapkan prinsip single responsibility

Pada contoh di bawah ini, kita memisahkan tanggung jawab `UserService` menjadi beberapa komponen yang berbeda, yaitu `Validator`, `UserRepository`, dan `EmailService`. Setiap komponen bertanggung jawab untuk melakukan satu hal saja. `UserService` bertanggung jawab untuk mengorkestrasi pendaftaran pengguna dengan menggunakan komponen yang berbeda.

1. Validator: Bertanggung jawab hanya untuk memvalidasi data pengguna.
2. UserRepository: Bertanggung jawab hanya untuk menyimpan data pengguna ke database.
3. EmailService: Bertanggung jawab hanya untuk mengirim email konfirmasi.
4. UserService: Bertanggung jawab untuk mengoordinasikan komponen-komponen di atas dalam proses pendaftaran pengguna.

```go
package main

import (
	"fmt"
	"net/smtp"
)

// User adalah entitas pengguna
type User struct {
	Name  string
	Email string
}

// Validator bertanggung jawab untuk memvalidasi data pengguna
type Validator struct{}

func (v *Validator) Validate(user User) error {
	if user.Name == "" || user.Email == "" {
		return fmt.Errorf("invalid user data")
	}
	return nil
}

// UserRepository bertanggung jawab untuk menyimpan pengguna ke database
type UserRepository struct{}

func (r *UserRepository) Save(user User) {
	// Simpan ke database (contoh sederhana)
	fmt.Println("Saving user to database...")
}

// EmailService bertanggung jawab untuk mengirim email konfirmasi
type EmailService struct{}

func (e *EmailService) SendConfirmation(user User) error {
	auth := smtp.PlainAuth("", "example@gmail.com", "password", "smtp.example.com")
	to := []string{user.Email}
	msg := []byte("Subject: Registration Successful!\n\nWelcome to our service!")
	err := smtp.SendMail("smtp.example.com:587", auth, "example@gmail.com", to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

// UserService mengorkestrasi pendaftaran pengguna dengan menggunakan komponen yang berbeda
type UserService struct {
	validator    *Validator
	repository   *UserRepository
	emailService *EmailService
}

func (s *UserService) Register(user User) error {
	// Validasi data pengguna
	if err := s.validator.Validate(user); err != nil {
		return err
	}

	// Simpan ke database
	s.repository.Save(user)

	// Kirim email konfirmasi
	if err := s.emailService.SendConfirmation(user); err != nil {
		return err
	}

	fmt.Println("User registered successfully!")
	return nil
}

func main() {
	validator := &Validator{}
	repository := &UserRepository{}
	emailService := &EmailService{}
	userService := &UserService{
		validator:    validator,
		repository:   repository,
		emailService: emailService,
	}

	user := User{Name: "John Doe", Email: "john.doe@example.com"}
	if err := userService.Register(user); err != nil {
		fmt.Println("Error:", err)
	}
}
```

## O - Open/Closed Principle

Tujuan utama dari principle ini adalah untuk menjaga agar code yang sudah ada tidak rusak ketika kita mengimplementasikan fitur baru.

Sebuah class dikatakan open jika kita dapat mengembangkannya, membuat subclass, dan melakukan apa pun yang kita inginkan dengan class tersebut-menambahkan method atau field baru, override perilaku dasar, dll. Beberapa bahasa pemrograman mengizinkan kita untuk membatasi perluasan lebih lanjut dari sebuah class dengan kata kunci khusus, seperti `final`. setelah itu, class tersebut tidak lagi open. Pada saat yang sama, class closed (bisa dikatakan selesai) jika sudah 100% siap untuk digunakan oleh class lain-interfacenya didefinisikan dengan jelas dan tidak akan diubah di masa depan.

Ketika pertama kali belajar mengenai principle ini, kita mungkin akan dibuat bingung karena kata "open" dan "closed" terdengar saling terpisah. Tetapi dalam principle ini, sebuah class dapat bersifat `Open` (untuk ekstensi) dan `Closed` (untuk modifikasi) pada saat yang sama.

Jika sebuah class sudah didevelop, diuji, direview, dan disertakan dalam suatu kerangka kerja atau digunakan dalam suatu aplikasi, mencoba mengotak-atik codenya akan sangat beresiko. Alih-alih mengubah code class secara langsung, kita dapat membuat subclass dan mengganti bagian dari class asli yang ingin kita ubah perilakunya. Kita akan mencapai tujuan yang sama tanpa merusak client yang sudah ada dari class asli.

Principle ini tidak dimaksudkan untuk diterapkan pada semua perubahan sebuah class. Jika kita mengetahui ada bug dalam class, langsung saja perbaiki, jangan membuat subclass untuk itu. Class anak tidak seharusnya bertanggung jawab atas masalah yang ada di class induknya.

## Example

- Contoh melanggar prinsip open/closed
  Dicontoh ini, kita akan menambahkan diskon dengan menambahkan logic langsung ke fungsi `CalculateTotalPrice`. Ini melanggar prinsip open/closed karena kita harus memodifikasi fungsi tersebut setiap kali ingin menambah atau mengubah jenis produk. </br>
  Jika nanti kita ingin menambah jenis produk baru, misalnya "clothing" dengan diskon 15%, kita harus mengubah kode dalam fungsi CalculateTotalPrice. Ini melanggar prinsip OCP.

```go
package main

import "fmt"

type Product struct {
    Name  string
    Price float64
    Type  string
}

func CalculateTotalPrice(products []Product) float64 {
    total := 0.0

    for _, product := range products {
        switch product.Type {
            case "electronic":
                // Diskon 10% untuk produk elektronik
                total += product.Price * 0.9
            case "groceries":
                // tidak ada diskon untuk produk kebutuhan pokok
                total += product.Price
            default:
                // tidak ada diskon untuk jenis produk lainnya
                total += product.Price
        }
    }

    return total
}

func main() {
	products := []Product{
		{"Laptop", 1000, "electronics"},
		{"Apple", 2, "groceries"},
	}

	total := CalculateTotalPrice(products)
	fmt.Printf("Total Price: %.2f\n", total)
}
```

- Contoh menerapkan prinsip open/closed
  Pada contoh ini, kita bisa menggunakan antarmuka (interface) yang memungkinkan kita menambahkan jenis produk baru tanpa mengubah kode yang ada.

```go
package main

import "fmt"

type IProduct interface {
	GetPrice() float64
}

// Product adalah struct untuk semua jenis produk
type Product struct {
	Name     string
	Price    float64
	Discount float64 // Diskon dalam persentase, misalnya 10% = 0.10
}

func NewProduct(name string, price, discount float64) IProduct {
	return &Product{
		Name:     name,
		Price:    price,
		Discount: discount,
	}
}

// GetPrice mengembalikan harga setelah diskon
func (p Product) GetPrice() float64 {
	return p.Price * (1 - p.Discount)
}

// CalculateTotalPrice menghitung total harga semua produk
func CalculateTotalPrice(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func main() {
	products := []Product{
		{Name: "Laptop", Price: 1000, Discount: 0.10}, // Elektronik, diskon 10%
		{Name: "Apple", Price: 2, Discount: 0.0},      // Sembako, tidak ada diskon
		{Name: "T-Shirt", Price: 20, Discount: 0.15},  // Pakaian, diskon 15%
	}

	total := CalculateTotalPrice(products)
	fmt.Printf("Total Price: %.2f\n", total)
}
```

## L - Liskov Substitution Principle

## I - Interface Segregation Principle

## D - Dependency Inversion Principle

## Creational Patterns

- [Abstract Factory](./creational-pattern/abstract-factory/README.md)
- [Builder](./creational-pattern/builder/README.md)
- [Factory Method](./creational-pattern/factory-method/README.md)
- [Prototype](./creational-pattern/prototype/README.md)
- [Singleton](./creational-pattern/singleton/README.md)

## Structural Patterns

- [Adapter](./structural-pattern/adapter/README.md)
- [Bridge](./structural-pattern/bridge/README.md)
- [Composite](./structural-pattern/composite/README.md)
- [Decorator](./structural-pattern/decorator/README.md)
- [Facade](./structural-pattern/facade/README.md)
- [Flyweight](./structural-pattern/flyweight/README.md)
- [Proxy](./structural-pattern/proxy/README.md)

## Behavioral Patterns

- [Chain of Responsibility](./behavioral-pattern/chain-of-responsibility/README.md)
- [Command](./behavioral-pattern/command/README.md)
- [Iterator](./behavioral-pattern/iterator/README.md)
- [Mediator](./behavioral-pattern/mediator/README.md)
- [Memento](./behavioral-pattern/memento/README.md)
- [Observer](./behavioral-pattern/observer/README.md)
- [State](./behavioral-pattern/state/README.md)
- [Strategy](./behavioral-pattern/strategy/README.md)
- [Template Method](./behavioral-pattern/template-method/README.md)
- [Visitor](./behavioral-pattern/visitor/README.md)
