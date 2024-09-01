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
