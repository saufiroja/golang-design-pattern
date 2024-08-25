# Visitor Pattern

Visitor pattern adalah salah satu design pattern yang memungkinkan kita untuk menambah operasi baru tanpa mengubah class-classnya. Ini dilakukan dengan cara mendefinisikan operasi dalam sebuah class terpisah yang disebut visitor dan mengizinkan object untuk menerima visitor tersebut. Visitor pattern sangat berguna ketika kita memiliki structure object yang kompleks dan ingin menjalankan operasi yang berbeda pada setiap element tanpa harus memorodifikasi element-element tersebut.

# Problem

Misalkan kita memiliki system manajemen dokumen dengan berbagai tipe dokumen, seperti PDF, Word, dan CSV. Setiap tipe dokumen memiliki cara yang berbeda untuk diolah, seperti diekspor, divalidasi, atau dianalisis. Jika kita menambahkan method baru ke setiap tipe dokumen, kita harus memodifikasi semua class dokumen yang ada, yang bisa menyebabkan masalah skalabilitas dan pemeliharaan.

# Solution

Dengan menggunakan visitor pattern, Kita bisa membuat 1 atau lebih class visitor yang mengimplementasikan operasi baru tanpa harus mengubah strukture dokumen. Setiap dokument akan memiliki method `accept` yang menerima visitor sebagai parameter. Visitor tersebut kemudian akan mengeksekusi operasi yang relevan tergantung pada tipe dokumen yang diterima.
