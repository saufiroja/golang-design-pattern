# Template Method Pattern

Template method pattern adalah salah satu design pattern dalam pemrograman yang memungkinkan kita mendefinisikan kerangka algorithm dalam method class induk, sementara subclass dapata mengisi atau mengganti beberapa langkah tertentu dari algorithm tersebut tanpa mengubah struktur dari algorithm secara keseluruhan.

# Problem

Misalnya kita memiliki system untuk mengirimkan OTP (One Time Password) melalui email dan SMS. Kedua proses ini memiliki langkah-langkah yang sama, yaitu:

1. Generate OTP
2. Save OTP
3. Prepare OTP message
4. Send OTP message

Namus, terdapat perbedaan pada cara mempersiapkan OTP message, koneksinya, dan cara mengirimkan OTP message. Jika kita tidak menggunakan template method pattern, mungkin akan ada banyak duplikasi kode untuk langkah-langkah yang sama, seperti generate OTP dan save OTP.

# Solution

Kita dapat menggunakan template method pattern untuk mengatasi masalah tersebut. Dengan template method pattern, kita dapat membuat class induk yang memiliki method-method yang sama untuk langkah-langkah yang sama, dan membiarkan subclass menentukan detail spesifik untuk setiap method pengiriman OTP.

# Kelebihan dan Kekurangan

1. Kelebihan

   - Kita dapat membiarkan client override hanya bagian tertentu dari algorithm, sehingga tidak terlalu terpengaruh oleh perubahan pada algorithm.
   - Dapat mengurangi duplikasi kode.

2. Kekurangan
   - Beberapa client mungkin dibatasi oleh kerangka algorithm yang disediakan.
   - Kita mungkin melanggar prinsip Liskov Substitution Principle dengan menekan implementasi langkah default melalui subclass.
   - Method template mungkin cenderung lebih sulit untuk dipertahankan, semakin banyak langkah yang mereka miliki.
