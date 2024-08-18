# Iterator Pattern

Iterator pattern adalah design pattern yang memungkinkan traversal elemen-elemen dari suatu kumpulan data collection secara satu per satu tanpa perlu mengetahui struktur internal dari kumpulan data tersebut. Dengan kata lain, iterator pattern memberikan cara yang seragam untuk mengakses elemen dari suatu collection, seperti array, map, or custom collection, tanpa harus mengetahu detail implementasi dari collection tersebut.

# Problem

Misalkan kita memiliki sebuah collection data yang terdiri dari berbagai jenis object, dan kita ingin mengakses setiap elemen dalam collection tersebut satu per satu. Namun, karna collection ini memiliki structure internal yang kompleks atau bervariasi, kita mungkin kesuiltan mengakses elemen-elemen tersebut secara konsisten dan efisien.

Misalnya, kita memiliki sebuah collection buku dalam perpustakaan yang di atur dalam berbagai format: beberapa dalam list, map dengan key, dan beberapa dalam tree structure berdasarkan abjad. Jika kita ingin mengakses semua buku ini secara konsisten, maka kita akan membutuhkan sebuah mekanisme yang dapat menangani berbagai jenis struktur ini tanpa harus mengubah code traversal setiap kali format collection berubah.

# Solusi

Dengan menerapkan iterator pattern, kita dapat membuat sebuah abstraksi yuang memungkinkan traversal dari elemen-elemen dalam collection ttanpa perlu peduli dengan detail implementasi dari collection tersebut. Iterator pattern akan menyediakan method seperti `next()`, `hasNext()`, `currentItem()` yang memungkinkan kita untuk menavigasi melalui elemen dalam collection dengan cara yang konsisten.

# Use Case

1. Traversal Data dari database
2. Proses stream data
3. Parsing dan Traversal data dari file

# Kekurangan dan Kelebihan

1. Kelebihan

   1. Single Responsibility Principle: Kita dapat membersihkan code client dan collection dengan mengekstrak traversal logic ke dalam class yang terpisah.
   2. Open/Closed Principle: Kita dapat menambahkan jenis collection baru dan meneruskannya ke code yang sudah ada tanpa harus mengubah code yang sudah ada.
   3. Kita dapat melakukan iterasi pada koleksi yang sama secara paralel karena setiap objek iterator berisi status iterasinya sendiri.
   4. Kita dapat menunda perulangan dan melanjutkannya saat dibutuhkan.

2. Kekurangan
   1. Menerapkan pola ini bisa jadi berlebihan jika aplikasi kita hanya bekerja dengan koleksi sederhana.
   2. Menggunakan iterator mungkin kurang efisien dibandingkan dengan menelusuri elemen-elemen dari beberapa koleksi khusus secara langsung.

# Implementation

```go

type Iterator[T any] interface {
	HasNext() bool
	Next() *T
}

// collection
type Collection[T any] interface {
	CreateIterator() Iterator[T]
}

// concrete iterator
type BookIterator struct {
	books []string
	index int
}

// concrete collection
type BookCollection struct {
	books []string
}

func (bc *BookCollection) CreateIterator() Iterator[string] {
	return &BookIterator{
		books: bc.books,
		index: 0,
	}
}

// concrete iterator
func (bi *BookIterator) HasNext() bool {
	return bi.index < len(bi.books)
}

func (bi *BookIterator) Next() *string {
	if bi.HasNext() {
		book := bi.books[bi.index]
		bi.index++
		return &book
	}

	return nil
}
```

# Contoh penggunaan

```go
func main() {
    books := []string{"book1", "book2", "book3", "book4", "book5"}

    bookCollection := &BookCollection{
        books: books,
    }

    iterator := bookCollection.CreateIterator()

    for iterator.HasNext() {
        book := iterator.Next()
        fmt.Println(*book)
    }
}
```
