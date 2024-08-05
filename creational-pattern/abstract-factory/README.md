# Abstract Factory

Abstract Factory adalah salah satu dari creational design pattern yang memberikan interface untuk membuat membuat family dari object terkait atau bergantung tanpa harus menentukan class spesifik. Dalam context pemrogramanan, pattern ini membantu dalam membuat system yang dapat dengan mudah diperluas dan dirubah tanpa harus mengubah code yang sudah ada.

# Konsep Abstract Factory

Konsep utama dari Abstract Factory adalah:

1. **Interface**: Mendefinisikan interface untuk membuat object, bukan instansiasi langsung class tertentu.
2. **Family of Product**: Membuat object yang saling berhubungan.
3. **Decoupling**: Memisahkan pembuatan object dengan implementasi class.

# Struktur Abstract Factory

1. **Abstract Factory**: Sebuah interface dengan mothod untuk membuat product abstrak.
2. **Concrete Factory**: Implementasi dari Abstract Factory yang membuat product konkret.
3. **Abstract Product**: Interface untuk product yang akan dibuat.
4. **Concrete Product**: Implementasi dari Abstract Product.
5. **Client**: Menggunakan Abstract Factory dan product yang dibuat.
