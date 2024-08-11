# Prototype Pattern

Prototype adalah design patern yang digunakan untuk membuat object baru dengan menyalin obect yang sudah ada. Pattern ini sangat berguna ketika proses pembuatan object sangat kompleks dan kamu perlu membuat banyak object yang mirip satu sama lain.

# Problem

Misalkan dalam aplikasi yang memiliki berbagai tipe user dengan sejumlah besar konfigurasi yang berbeda. Membuat object untuk setiap tipe use secara manual akan sangat memakan waktu dan rentan terhadap kesalahan, terutama jika konfigurasi tersebut sering berubah.

Proses ini akan memerlukan banyak duplikasi code dan konsumsi memory yang tidak efisien. Menambahkan atau mengubah setelan user uga akan membutuhkan perubahan yang signifikan diberbagai tempat dalam code, yang meningkatkan risiko kesalahan.

# Solution

Menggunakan prototype pattern memungkinkan kamu membuat object baru dengan cepat tanpa harus memulai dari awal. Kamu dapat membuat obect baru dengan menyalin object yang sudah ada dan mengubah bagian yang diperlukan.

Prototype pattern mengurangi kebutuhan untuk menulis code yang sama berulang-ulang untuk membuat object dengan properti serupa, sehingga membuat code lebih bersih dan lebih mudah dikelola.

Jika ada perubahan dalam structure atau properti object, kamu hanya perlu mengubah prototype object, dan perubahan tersebut akan otomatis diterapkan pada semua object yang dibuat dari prototype tersebut.

# Structure

Prototype pattern di Go biasanya melibatkan structure yang berisi data state dan method untuk mengclone object tersebut.

```go

// prototype interface
type UserPrototype interface {
    Clone() UserPrototype
}

// concrete prototype
type User struct {
    Name  string
    Email string
}

// method to clone object
func (u *User) Clone() UserPrototype {
    // shallow copy
    return &User{
        Name:  u.Name,
        Email: u.Email,
    }
}

func main() {
    user1 := &User{
        Name:  "John Doe",
        Email: "test@mail.com",
    }

    user2 := user1.Clone()

    fmt.Println(user1)
    fmt.Println(user2)
}

```

# Kapan Harus Menggunakan Prototype Pattern

# Kelebihan dan Kekurangan

1. Kelebihan

   - Kamu dapa mengclone objject tanpa mengaitkan ke class konkrit
   - Kamu dapat menyingkirkan code inisialisasi yang berulang demi mengclone prototype yang telah dibuat sebelumnya
   - Kamu dapat menghasilkan object yang kompleks dengan mudah
   - Kamu mendapatkan alternatif untuk inheritance ketika berurusan dengan preset configuration untuk object yang kompleks

2. Kekurangan
   - Mengclone object yang kompleks yang memiliki referensi melingkar mungkin sangat sulit
