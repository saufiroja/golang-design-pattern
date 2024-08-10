# Abstract Factory

Abstract Factory adalah salah satu dari creational design pattern, pattern ini digunakan untuk membuat interface atau class abstrak untuk membuat kelompok object yang terkait tanpa harus menentukan class konkretnya. Abstrac factory memungkinkan code untuk lebih fleksibel dan mudah diubah karena tidak tergantung pada class konkrit.

# Problem

Misalkan kamu sedang membuat aplikasi payment service dengan berbagai jenis pembayaran seperti Paypal, Stripe, dan lain-lain. Setiap Jenis pembayaran ini memiliki cara implementasi yang berbeda-beda(misalnya, cara pembayaran, cara verifikasi, dll). Jika kamu membuat obect dari masing-masing enis pembayaran secara langsung, maka ini akan membuat code menjadi sulit diubah dan tidak fleksibel. Jika kamu ingin menambahkan jenis pembayaran baru atau mengubah implementasi dari salah satu method yang sudah ada, maka kamu harus melakukan perubahan pada banyak tempat, yang beresiko membuat code menjadi error.

# Solusi

Dengan menggunakan abstract factory, kamu bisa membuat 1 interface untuk membuat kelompok object terkait, seperti `PaymentFactory`. Setiap enis pembayaran akan memiliki implementasinya sendiri, misalnya `PaypalFactory`, `StripeFactory`, dll. Ketika aplikasi memerlukan pembayaran, ia akan menggunakan factory yang tepat tanpa harus mengetahui detail implementasinya dari masing-masing jenis pembayaran. Ini memungkinkan kamu untuk menambah atau mengubah jenis pembayaran tanpa mengubah banyak tempat pada code.

# Struktur Abstract Factory

1. **Abstract Factory**: Sebuah interface dengan mothod untuk membuat product abstrak.
2. **Concrete Factory**: Implementasi dari Abstract Factory yang membuat product konkret.
3. **Abstract Product**: Interface untuk product yang akan dibuat.
4. **Concrete Product**: Implementasi dari Abstract Product.
5. **Client**: Menggunakan Abstract Factory dan product yang dibuat.
