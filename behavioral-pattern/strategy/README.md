# Strategy Pattern

Strategy pattern adalah salah satu design pattern yang termasuk dalam kategori behavioral pattern. Pattern ini digunakan untuk mendefinisikan sebuah keluarga algoritma, mengenkapsulasi masing-masing algoritma ke dalam class yang berbeda, dan memungkinkan alogirtma-algoritma tersebut untuk saling dipertukarkan tanpa mengubah code client yang menggunakannya. Dengan kata lain, strategy pattern memungkinkan kita untuk memiliki algoritma yang sesuai pada saat runtime.

# Problem

Misalkan Kita memiliki sebuah aplikasi e-commerce yang menawarkan berbagai method pembayaran (misalnya: transfer bank, kartu kredit, e-wallet, dll). Jika kita mencoba menangani semua logic pembayaran dalam 1 class, code akan menjadi kompleks dan sulit untuk dipelihara. Setiap kali ada method pembayaran baru yang ingin ditambahkan, kita harus modify class tersebut, yang melanggar prinsip Open/Closed Principle dalam SOLID (sebuah class seharusnya `open for extension, but closed for modification`).

# Solution

Dengan menggunakan strategy pattern, kita dapat membuat interface `PaymentStrategy` yang memiliki method `Pay(amount float64)`. Setiap method pembayaran kemudian akan diimplementasikan dalam class terpisah, seperti `CreditCardPayment`, `BankTransferPayment`, `EwalletPayment`, dll, yang semuanya mengimplementasikan interface `PaymentStrategy`.
Ini menunukkan kita untuk dengan mudah menambahkan method pembayaran baru tanpa harus mengubah code client yang menggunakannya. Kita hanya perlu membuat class baru yang mengimplementasikan interface `PaymentStrategy`.

# Use Case

1. Pemilihan Algoritma Kompresi: Misalnya, Kita mungkin memiliki beberapa algoritma kompresi yang berbeda (misalnya: zip, rar, 7z, dll) dan kita ingin memungkinkan pengguna untuk memilih algoritma kompresi yang sesuai pada saat runtime.
2. Authentication Multi-Provider: Jika aplikasi kita mendukung berbagai metode autentikasi (misalnya: JWT, OAuth, Basic Auth, dll), kita dapat menggunakan strategy pattern untuk memungkinkan pengguna memilih metode autentikasi yang sesuai pada saat runtime.
3. Load Balancer: Misalnya, kita mungkin memiliki beberapa algoritma load balancing yang berbeda (misalnya: round-robin, least-connection, random, dll) dan kita ingin memungkinkan pengguna untuk memilih algoritma load balancing yang sesuai pada saat runtime.

# Kelebihan dan Kekurangan

1. **Kelebihan**
   - Kamu bisa menukar algoritma yang digunakan pada saat runtime.
   - Kita dapat mengisolasi detail implementasi algoritma dari code yang menggunakannya.
   - Kita dapat mengganti inheritance dengan composition.
   - Open/Closed Principle: Kita dapat menambahkan algoritma baru tanpa harus mengubah code yang sudah ada.
2. **Kekurangan**
   - Jika memiliki code yang jarang berubah, tidak ada alasan untuk menggunakan strategy pattern.
   - Client harus mengetahu perbedaan antar strategy untuk memilih yang sesuai.

# Kapan Menggunakan Strategy Pattern

1. Kita memiliki beberapa algoritma atau perilaku yang berbeda-beda dan ingin membuatnya dapat dipertukarkan secara dinamis.
2. Kita ingin menghindari code yang penuh dengan kondisi if-else.
3. Kita ingin membuat code yang mudah untuk dipelihara dan diperluas tanpa harus mengubah code yang sudah ada.
