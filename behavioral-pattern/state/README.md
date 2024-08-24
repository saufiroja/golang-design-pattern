# State Pattern

State pattern adalah salah satu design pattern yang digunakan untuk mengizinkan object mengubah perilakunya ketika status internalnya berubah. pattern ini sangat berguna ketika suatu object harus mengubah perilakunya tergantung pada status internal yang dimiliki

# Problem

Bayangkan kita sedang mengembangkan aplikasi yang mengelola pemesanan tiket pesawat. Setiap pemesanan tiket bisa berada dalam beberapa state, seperti:

- **Baru dibuat(NEW)**: Pemesanan tiket baru dibuat
- **Dibayar(PAID)**: Pemesanan tiket sudah dibayar
- **Dibatalkan(CANCELLED)**: Pemesanan tiket dibatalkan
- **Selesai(COMPLETED)**: Pemesanan tiket selesai

Tanpa state pattern, kita mungkin akan memiliki code yang penuh dengan if-else atau switch case untuk menangani berbagai state yang ada. Hal ini bisa menadi sulit di pelihara, terutama ketika state baru ditambahkan atau diubah.

# Solusi

Dengan state pattern, kita dapat memisahkan logika untuk setiap state ke dalam class yang berbeda. Setia state akan memiliki class sendiri, dan setiap class akan mengimplementasikan interface yang sama. Dengan cara ini, code menjadi lebih modular, mudah di pelihara, dan lebih fleksibel.

# Use Case

1. **Workflow Engine**: Menangani berbagai status dalam alur kerja seperti state persetujuan, penolakan, pengiriman, dll.
2. **Game Development**: Menangani berbagai state dalam game seperti state bermain, pause, game over, dll.
3. **Order Management System**: Menangani berbagai state dalam pemesanan seperti state order placed, payment pending, shipped, delivered, cancelled, dll.

# Kapan harus menggunakan State Pattern?

1. Object memiliki beberapa state yang berbeda, dan state tersebut mempengaruhi perilaku object.
2. Code mulai dipenuhi dengan banyak if-else atau switch case untuk menangani berbagai state.
3. Dibutuhkan solusi yang modular, mudah di pelihara, dan fleksibel.

# Kelebihan dan Kekurangan

1. **Kelebihaan**
   - Single Responsibility Principle: Mengatur code yang terkait dengan status tertentu ke dalam class yang terpisah.
   - Open/Closed Principle: Memperkenalkan state baru tanpa mengubah class state yang sudah ada atau contextnya.
   - Menyederhanakan code context dengan menghilangkan if-else atau switch case yang berlebihan.
2. **Kekurangan**
   - Menerapkan pattern ini bisa adi berlebihan jika state machine hanya memiliki sedikit state atau jarang berubah.

# Implementasi

- State Interface

```go
// state interface
type OrderState interface {
	Next(*Order) error
	Cancel(*Order) error
	Status() string
}

// order structure
type Order struct {
	state OrderState
}

// method to change state
func (o *Order) ChangeState(state OrderState) {
	o.state = state
}

// next state transition
func (o *Order) Next() error {
	return o.state.Next(o)
}

// cancel state transition
func (o *Order) Cancel() error {
	return o.state.Cancel(o)
}

// get current state
func (o *Order) Status() string {
	return o.state.Status()
}
```

- State New (Baru Dibuat)

```go
type NewOrder struct{}

func (o *NewOrder) Next(order *Order) error {
	order.ChangeState(&PaidState{})
	return nil
}

func (o *NewOrder) Cancel(order *Order) error {
	order.ChangeState(&CancelledState{})
	return nil
}

func (o *NewOrder) Status() string {
	return "order baru dibuat"
}
```

- State Paid (Dibayar)

```go
type PaidState struct{}

func (o *PaidState) Next(order *Order) error {
	order.ChangeState(&CompletedState{})
	return nil
}

func (o *PaidState) Cancel(order *Order) error {
	order.ChangeState(&CancelledState{})
	return nil
}

func (o *PaidState) Status() string {
	return "order sudah dibayar"
}
```

- State Completed (Selesai)

```go
type CompletedState struct{}

func (o *CompletedState) Next(order *Order) error {
	return fmt.Errorf("order sudah selesai")
}

func (o *CompletedState) Cancel(order *Order) error {
	return fmt.Errorf("order sudah selesai dan tidak bisa dibatalkan")
}

func (o *CompletedState) Status() string {
	return "order sudah selesai"
}
```

- State Cancelled (Dibatalkan)

```go
type CancelledState struct{}

func (o *CancelledState) Next(order *Order) error {
	return fmt.Errorf("order sudah dibatalkan")
}

func (o *CancelledState) Cancel(order *Order) error {
	return fmt.Errorf("order sudah dibatalkan")
}

func (o *CancelledState) Status() string {
	return "order sudah dibatalkan"
}
```

- Main Program

```go
func main() {
	order := &Order{state: &NewOrder{}}

	fmt.Println(order.Status()) // "Order baru dibuat."

	order.Next()
	fmt.Println(order.Status()) // "Order sudah dibayar."

	order.Next()
	fmt.Println(order.Status()) // "Order selesai."

	err := order.Cancel()
	if err != nil {
		fmt.Println(err) // "Order sudah selesai dan tidak bisa dibatalkan."
	}
}
```
