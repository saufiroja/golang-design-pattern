# Factory Method

Factory method adalah salah satu design pattern dalam kategori creational pattern. Fungsi dari factory method adalah untuk mengizinkan sebuah class atau object untuk membuat object lain. Tujuan utamanya adalah untuk menyediakan sebuah interface bagi pembuatan object, tetapi memungkinkan subclass untuk mengubah tipe object yang akan di buat.

Pattern ini memisahkan proses pembuatan object dari penggunaan object tersebut., memberikan fleksibilitas untuk mengubah implementasi yang digunakan untuk membuat object tanpa mengubah code yang menggunakan object tersebut. Dengan menggunakan factory method, code client dapat berinteraksi dengan object melalui interface yang umum tanpa perlu mengetahui implementasi spesifik dari object yang digunakan.

# Komponen utama factory method

- Product: Interface yang mendefinisikan object yang akan di buat oleh factory method.
- ConcreteProduct: Implementasi dari interface Product.
- Creator: Class yang mendeklarasikan factory method yang mengembalikan object dari type Product. Class ini biasanya memiliki implementasi default dari factory method yang dapat diubah oleh subclass.
- ConcreteCreator: Subclass yang mengimplementasikan factory method untuk membuat dan mengembalikan objject dari type product tertentu.

# Contoh kasus

Factory method bisa sangat berguna dalam berbagai situasi. Berikut adalah contoh penggunaan factory method:

1. System Payment
   - Mengelola berbagai jenis pembayaran seperti kartu kredit, transfer bank, dan e-wallet dengan mengizinkan code client untuk menggunakan method pembayaran secara transparan.
2. System Database Connection
   - Menghubungkan berbagai jenis database seperti MySQL, PostgreSQL, dan MongoDB dengan mengizinkan code client untuk menggunakan database
3. Notification System
   - Mengirimkan notifikasi melalui berbagai media seperti email, SMS, dan push notification dengan mengizinkan code client untuk mengirimkan notifikasi tanpa perlu mengetahui media yang digunakan.
4. System Logging
   - Mencatat log aplikasi ke berbagai media seperti file, database, dan console dengan mengizinkan code client untuk mencatat log tanpa perlu mengetahui media yang digunakan.
5. System Shipping
   - Mengatur pengiriman barang ke berbagai lokasi dengan mengizinkan code client untuk mengatur pengiriman tanpa perlu mengetahui lokasi pengiriman.

# Kelebihan dan Kekurangan

1. Kelebihan

   - Kamu menghindari duplikasi code dengan memisahkan creator dan concrete product.
   - Single Responsibility Principle: Kamuu bisa memindahkan product creation ke 1 tempat kedalam program, sehingga code lebih mudah untuk di support.
   - Open/Closed Principle: Kamu dapat memperkenalkan product type baru kedaam program tanpa merusak code client yang sudah ada.

2. Kekurangan
   - Code dapat menjadi lebih rumit karena anda perlu memperkenalkan banyak subclass baru untuk mengimplementasikan pattern ini. Skenario kasus terbaik adalah ketika anda memperkenalkan pattern ke dalam hirarki creator class yang sudah ada.
