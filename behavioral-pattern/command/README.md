# Command Pattern

Command Pattern adalah salah satu design pattern yang termasuk dalam kategori behavioral pattern. Pattern ini digunakan untuk mengenkapsulasi request sebagai object, sehingga memungkinkan kamu untuk memparameterisasi object dengan operasi yang berbeda, menjadwalkan operasi atau meletakkan operasi dalam antrian.

# Problem

Masalah yang sering muncul dalam pengembangan software adalah ketika kita perlu mengimplementasikan operasi yang mungkin harus dijalankan secara terpisah, mungkin diundur, atau bahkan dibatalkan. Misalnya, dalam aplikasi seperti e-commerce, kita mungkin perlu mendukung operasi seperti `Undo Purchase` atau `Redo Purchase`. Jika setiap tindakan dalam system diimplementasikan sebagai fungsi atau method biasa, manajemen undo/redo dan fleksibilitas lainnya akan menjadi sangat sulit.

# Solution

Command patttern memecahkan masalah ini dengan memisahkan request dari penerima yang melakukan request tersebut. Dalam hal ini, setiap operasi seperti `Purchase`, `Cancel Purcahse`, atau `Undo Purchase` diimplementasikan sebagai class command yang mengenkapsulasi semua informasi yang diperlukan untuk melakukan operasi tersebut. Dengan begitu, operasi ini dapat disimpan, dikelola, atau digabungkan dengan lebih mudah.

# Use Case

- Transaksi pada e-commerce
  Setiap transaksi, seperti `Purchase`, `Cancel Purchase`, atau `Undo Purchase` dapat diimplementasikan sebagai class command yang berbeda.

# Kekurangan dan Kelebihan

1. Kelebihan
   - Single Responsibility Principle: Kitta bisa memisahkan class yang memanggil operasi dari class yang menjalankan operasi tersebut.
   - Open/Closed Principle: Kita dapat memperkenalkan perintah baru kedalam aplikasi tanpa merusah code client yang sudah ada.
   - Kita dapat mengimplementasikan operasi seperti `Undo` atau `Redo` dengan mudah.
   - Kita dapat mengimplementasikan eksekusi operasi
   - Kita dapat mengumpulkan serangkaian command sederhana menjadi perintah yang kompleks.
2. Kekurangan
   - Code dapat menjadi lebih rumit karena kita memperkanalkan banyak class baru antara pengirim dan penerima.
