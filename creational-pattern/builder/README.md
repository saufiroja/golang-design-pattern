# Builder Pattern

Builder pattern adalah salah satu dari creational design pattern yang bertujuan untuk memisahkan konstruksi object yang kompleks dari representasinya. Dengan menggunakan builder pattern, kita bisa membuat objject dengan banyak parameter tanpa harus menggunakan contrustor dengan banyak argumen, yang dapa membuat code sulit dibaca dan sulit dipahami.

# Problem

Misalkan kamu memiliki object `Person` yang memiliki banyak atribut seperti `name`, `age`, `gender`, `address`, `phone`, `email`, dan lain-lain. Jika kita ingin membuat object `Person` dengan atribut yang berbeda, kamu mungkin harus menggunakan banyak constructor yang kelebihan beban atau membuat object dengan menggunakan setter method. Ini akan membuat code sulit untuk dibaca dan dipahami.

```go
type Person struct {
    name string
    age int
    gender string
    address string
    phone string
    email string
}

func main() {
    // Create person with all attributes
    person1 := Person{
        name: "John Doe",
        age: 30,
        gender: "male",
        address: "Jl. Jendral Sudirman No. 1",
        phone: "08123456789",
        email: "test@gmail.com",
    }

    fmt.Println(person1)
}
```

# Solution

Dengan menggunakan builder pattern, kita dapat membuat object `Person` dengan cara yang lebih terstruktur dan terorganisir. Builder pattern memecah pembuatan object yang kompleks menjadi langkah-langkah yang lebih kecil dan terstruktur.

```go
type Person struct {
    name string
    age int
    gender string
    address string
    phone string
    email string
}

type PersonBuilder struct {
    person Person
}

func NewPersonBuilder() *PersonBuilder {
    return &PersonBuilder{}
}

func (pb *PersonBuilder) SetName(name string) *PersonBuilder {
    pb.person.name = name
    return pb
}

func (pb *PersonBuilder) SetAge(age int) *PersonBuilder {
    pb.person.age = age
    return pb
}

func (pb *PersonBuilder) SetGender(gender string) *PersonBuilder {
    pb.person.gender = gender
    return pb
}

func (pb *PersonBuilder) SetAddress(address string) *PersonBuilder {
    pb.person.address = address
    return pb
}

func (pb *PersonBuilder) SetPhone(phone string) *PersonBuilder {
    pb.person.phone = phone
    return pb
}

func (pb *PersonBuilder) SetEmail(email string) *PersonBuilder {
    pb.person.email = email
    return pb
}

func (pb *PersonBuilder) Build() Person {
    return pb.person
}

func main() {
    // Create person with all attributes
    person1 := NewPersonBuilder().
        SetName("John Doe").
        SetAge(30).
        SetGender("Male").
        SetAddress("Jl. Jendral Sudirman No. 1").
        SetPhone("08123456789").
        SetEmail("test@gmail.com").
        Build()

    fmt.Println(person1)

    // Create person with only name and age
    person2 := NewPersonBuilder().
        SetName("Jane Doe").
        SetAge(25).
        Build()

    fmt.Println(person2)
}
```

# Structure

- **Builder**: Interface yang mendefinisikan langkah-langkah konstruksi produk yang umum untuk semua jenis builder.
- **ConcreteBuilder**: Implementasi yang berbeda dari langkah-langkah konstruksi. ConcreteBuilder dapat menghasilkan product yang tidak mengikuti interface builder.
- **Product**: Object yang dihasilkan. Product yang di bangun oleh builder yang berbeda tidak harus memiliki hirarki class atau interface yang sama.
- **Director**: Mendefinisikan urutan langkah-langkah konstruksi, sehingga dapat membuat dan menggunakan kembali konfigurasi produk tertentu.
- **Client**: Membuat object builder dan mengatur director untuk membuat object.

# Kekurangan dan Kelebihan'

1. Kelebihan

   - Dapat membuat obect langkah demi langkah, menunda langka konstruksi, atau menjalankan langkah secara rekursif.
   - Dapat menggunakan kembali code konstruksi yang sama ketika membuat berbagai representasi produk.
   - Single Responsibility Principle. Anda dapat memisahkan konstruksi kompleks dari representasi produk.

2. Kekurangan
   - Keseluruhan code akan lebih rumit karena membutuhkan banyak class baru.
