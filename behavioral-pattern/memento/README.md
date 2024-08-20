# Memento Pattern

Memento adalah salah satu design pattern yang termasuk dalam kategori behavioral pattern. Pattern ini digunakan untuk menyimpan dan memulihkan keadaan atau state dari sebuah object tanpa melanggar prinsip encapsulation. Dengan kata lain, memento memungkinkan penyimpanan snapshot dari state object tertentu sehingga bisa dikembalikan ke state sebelumnya kapan saja.

# Problem

Bayangkan kamu sedang membuat sebuah aplikasi editor teks di mana pengguna bisa melakukan undo dan redo terhadap perubahan yang dilakukan. Setiap kali pengguna mengetik sesuatu, menambahkan format, atau mengahpus teks, state dari dokumen berubah. Problemnya adalh bagaiman caranya menyimpan state-state ini sehingga pengguna dapat dengan mudah kembali ke state sebelumnya tanpa merusak prinsip encapsulation dari object dokumen itu sendiri.

# Solution

Solusi untuk problem ini adalah dengan menggunakan memento pattern. Dengan menggunakan pattern ini, kita bisa membuat snapshot dari state object dokumen setiap kali ada perubahan, dan menyimpannya di dalam stack. Ketika pengguna ingin melakukan undo, kita bisa mengembalikan dokumen ke state sebelumnya dengan mengambil memento dari stack. Ini memungkinkan kita untuk mengimplementasikan undo dan redo dengan mudah tanpa mengubah struktur internal dari object dokumen itu sendiri.

# Use Case

1. System Undo/Redo: Misalnya dalam aplikasi pengelolaan konten atau editor online.
2. Simpan dan restore konfigurasi: Pada system yang mengizinkan pengguna untuk menyimpan dan mengembalikan konfigurasi sebelumnya.
3. Versioning Data: Ketika perlu menyimpan versi sebelumnya dari data dalam suatu aplikasi.

# Kekurangan dan Kelebihan

1. Kelebihan:
   - Kita dapat membuat snapshot dari state object tanpa melanggar prinsip encapsulation.
   - Kita dapat menyederhanakan originator code dengan membiarkan caretaker yang bertanggung jawab untuk menyimpan dan mengembalikan memento.
2. Kekurangan:
   - Aplikasi ini mungkin menghabiskan banyak RAM ika client terlalu sering membuat memento.
   - Caretakers harus melacak siklus originator agar dapat destroy memento yang tidak diperlukan.
   - Sebagian besa bahasa pemrograman seperti PHP, Python dan Javascript, tidak dapat menjamin bahwa state dalam memento tetap tidak tersentuh.

# Struktur

1. Originator: Object yang memiliki state yang ingin disimpan.
2. Memento: Object yang digunakan untuk menyimpan state dari originator. Ini biasanya adalah struct sederhana yang hanya berisi data state.
3. Caretaker: Object yang bertanggung jawab untuk menyimpan dan mengembalikan memento. Ini biasanya mengandung stack atau list untuk menyimpan memento.

# Implementation

```go
// Memento
type Memento struct {
	level int
	score int
}


// Originator
type Originator interface {
	Save() *Memento
	Restore(*Memento)
}

type GameWorld struct {
	level int
	score int
}

func (g *GameWorld) Save() *Memento {
	return &Memento{level: g.level, score: g.score}
}

func (g *GameWorld) Restore(m *Memento) {
	g.level = m.level
	g.score = m.score
}

// Caretaker
type SaveFileManager struct {
	memento []*Memento
}

func (s *SaveFileManager) Save(m *Memento) {
	s.memento = append(s.memento, m)
}

func (s *SaveFileManager) Load(index int) *Memento {
	return s.memento[index]
}
```

```go
func main() {
	gameWorld := GameWorld{level: 1, score: 0}
	saveFileManager := &SaveFileManager{}

	// save the initial state of the game
	saveFileManager.Save(gameWorld.Save())

	// increase the level and score
	gameWorld.level++
	gameWorld.score += 1000

	// save the new state of the game
	saveFileManager.Save(gameWorld.Save())

	// restore the initial state of the game
	gameWorld.Restore(saveFileManager.Load(0))

	fmt.Printf("Level: %d, Score: %d\n", gameWorld.level, gameWorld.score)
}
```
