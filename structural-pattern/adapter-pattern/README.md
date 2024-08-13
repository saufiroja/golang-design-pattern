# Adapter Pattern

Adapter pattern adalah salah satu dari design pattern dalam kategori structural pattern. Pattern ini memungkinkan 2 interface yang tidak kompatibel untuk bekerja sama dengan menghubungkan mereka melalui adapter class. Pattern ini bertindak sebagai perantara antara 2 interface yang berbeda, sehingga class tersebut bisa saling berinteraksi tanpa perlu mengubah code yang sudah ada.

# Problem

Bayangkan kita memiliki sebuah system lama yang sudah berfungsi dengan baik namun menggunakan interface atau method tertentu. Seiring waktu, kita mungkin ingin mengguanakan komponen baru yang lebih modern, namu komponen baru ini memiliki interface yang berbeda. Mengubah system lama untuk support interface baru ini bisa sangat mahal dan beresiko, terutama jika system lama digunakan diberbagai tempat.

## Contoh masalah

1. Sebuah aplikasi lama yang menggunakan library logging yang memiliki interface tertentu, tetapi kita ingin mengintegrasikan library logging baru yang memiliki interface yang berbeda.
2. Kita memiliki system yang berkomunikasi dengan API eksternal, namu api tersebut telah diperbarui dan menggunakan format atau protokol yang berbeda.

# Solution

Adapter pattern memungkinkan kita membuat class adapter yang menerjemahkan interface baru ke dalam format yang bisa dimengerti oleh system lama, atau sebaliknya.. Dengan cara ini, kita menggunakan komponen baru tanpa harus mengubah code yang sudah ada.

## Solusi umum

1. Membuat sebuah class adapter yang mengimplementasikan interface yang diinginkan dan didalamnya memanggil atau memetakan fungsi dari class atau komponen lain yang memiliki interface yang berbeda.
2. Adapter ini menjadi perantara yang mengubah bahasa atau format antara 2 class yang berbeda, sehingga mereka bisa berinteraksi.

# Use Cases

1. Interfacing dengan database:
   Jika kita mengubah database yang digunakan, kita bisa menggunakan adapter untuk memasttikan interface fungsi database tetap konsisten.
2. Integrasi API Eksternal:
   Saat kita beralih dari 1 service eksternal ke service lain dengan interface API yang berbeda, adapter bisa digunakan untuk menyembunyikan perbedaan dari code.
3. Logger Abstraction:
   Menggunakan adapter untuk mengabstraksi implementasi logging, sehingga jika kamu mengganti library logging, perubahan hanya dilakukan di adapter tanpa mengubah seluruh code aplikasi.

# Structure

- Target adalah antarmuka yang diharapkan oleh sistem.
- Adaptee adalah kelas dengan antarmuka berbeda yang ingin Anda gunakan.
- Adapter adalah kelas yang mengimplementasikan antarmuka Target dan memetakan fungsionalitas Adaptee ke dalam format yang bisa diterima oleh Target.

# Kapan Harus Menggunakan Adapter Pattern

- kita ingin menggunakan kelas atau komponen yang ada tetapi antarmukanya tidak kompatibel dengan kode yang ada.
- Kita ingin meminimalkan perubahan pada kode lama saat mengintegrasikan teknologi baru.
- Kita perlu mendukung antarmuka lama sambil beralih ke antarmuka baru secara bertahap.

# Kekurangan dan Kelebihan

1. Kelebihan

   - Single Responsibility Principle, Kita dapat memisahkan interface atau code konversi data dari logic bisnis utama program
   - Open/Closed Principle, Kita dapat memperkenalkan jenis adapter baru ke dalam program tanpa merusak code client yang ada

2. Kekurangan
   - Keseluruhan kompleksitas kode meningkat karena Anda perlu memperkenalkan serangkaian antarmuka dan kelas baru. Kadang-kadang lebih sederhana hanya dengan mengubah kelas layanan sehingga cocok dengan kode Anda yang lain.
