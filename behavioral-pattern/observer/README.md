# Observer Pattern

Observer pattern adalah salah satu dari design pattern yang digunakan untuk mengatur system dimana object-object berinteraksi satu sama lain. Pattern ini memungkinkan satu object atau disebut subject memberi tahu object lain atau disebut observer tentang perubahan status tanpa perlu mengikat satu sama lain secara ketat.
Observer pattern memungkinkan object bertindak sebagai subject untuk memiliki daftar Observer yang tertarik pada perubahan status. Ketika ada perubahan dalam subject, semua observer akkan diberitahu atau diupdate secara otomatis.. Pattern ini sering digunakan untuk mengimplementasikan systen event-driven atau notfikasi.

# Problem

Bayangkan sebuah system notification dimana berbagai komponen aplikasi perlu diberitahu ketika pengguna mengubah preferensi pengaturan mereka. Misalnya, saat pengguna mengubah pengaturan bahasa, interface pengguna, logging, dan komponen lain yang bergantung pad abahasa harus diupdate.
Dalam skenario ini, tanpa observer pattern, setiap komponen yang perlu di update harus secara langsung dipanggil dari subject, yang menyebabkan ketergantungan tingga dan code yang sulit untuk di maintain.

# Solution

Dengan menggunakan observer pattern, kita dapat membuat komponen-komponen tersebut menjadi observer dari subject. Ketika pengaturan pengguna berubah, subject cukup memanggil method notifikasi, dan semua observer akan di update secara otomatis tanpa perlu mengikat satu sama lain secara langsung. Ini membuat system lebih fleksibel dan modular.

# Kelebihan dan Kekurangan

1. Kelebihan
2. Kekurangan

# Use Case

1. Event-Driven System -> system dimana tindakan tertentu harus memicu berbagai reaksi di berbagai bagian system.
2. Logging system -> Ketika berbagai bagian aplikasi perlu mencatat log berdasarkan tindakan tertentu.
3. Notification system -> Untuk mengirimkan notifikasi keberbagai sub-system ketika event tertentu terjadi.
4. Data Binding -> Menghubungkan data antar berbagai layer aplikasi yang membutuhkan sinkronisasi data.

# Implementation

```go
// observer interface
type Observer interface {
	Update(event string)
}

// concrete observer A
type EmailNotification struct{}

func (e *EmailNotification) Update(event string) {
	fmt.Println("EmailNotifier received event:", event)
}

// concrete observer B
type SMSNotification struct{}

func (s *SMSNotification) Update(event string) {
	fmt.Println("SMSNotifier received event:", event)
}

// subject interface
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(event string)
}

// concrete subject
type EventManager struct {
	observers []Observer
}

func (e *EventManager) Register(observer Observer) {
	e.observers = append(e.observers, observer)
}

func (e *EventManager) Unregister(observer Observer) {
	for i, o := range e.observers {
		if o == observer {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
			break
		}
	}
}

func (e *EventManager) Notify(event string) {
	for _, observer := range e.observers {
		observer.Update(event)
	}
}
```

```go
func main() {
    eventManager := &EventManager{}
    emailNotifier := &EmailNotification{}
    smsNotifier := &SMSNotification{}

    eventManager.Register(emailNotifier)
    eventManager.Register(smsNotifier)

    eventManager.Notify("User logged in")
}
```
