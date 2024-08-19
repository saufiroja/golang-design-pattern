# Mediator Pattern

Mediator Pattern adalah design pattern yang digunakan untuk mengurangi kompleksitas komunikasi antar object dalam sebuah system dengan memperkenalkan sebuah object mediator yang bertindak sebagai penghubung. Object dalam system tidak saling berkomunikasi secara langsung, melainkan melalui object mediator, yang mengatur interaksi antar object. Pattern ini membantu dalam menjaga loose coupling antar object dalam system.

# Problem

Dalam aplikasi e-commerce, ada beberapa module seperti order, inventory, dan payment. Setiap module ini perlu berkomunikasi satu sama lain. Misalnya, ketika sebuah order dibuat, module order perlu memberitahu module inventory untuk mengurangi stock dan kemudian memperitahu module payment untuk melakukan pembayaran. Jika komunikasi dilakukan secara langsung antar module, maka akan teradi tight coupling antar module, yang membuat system sulit untuk di maintain.

# Solution

Dengan menerapkan mediator pattern, kita bisa membuat sebuah mediator yang mengatur komunikasi antara module order, inventory, dan payment. Setiap module hanya perlu mengetahui mediator, dan mediator yang akan mengatur interaksi antar module. Ini mengurangi ketergantungan langsung antar module, sehingga system lebih mudah untuk di maintain.

# Use Case

Mediator pattern bisa digunakan dalam konteks CQRS(Command Query Responsibility Segregation) dan ketika ada kebutuhan mengatur interaksi kompleks antara berbagai handler atau service dalam aplikasi. Di Go, Mediator biasanya digunakan bersama dengan pattern lainnya seperti Command Pattern dan Event Sourcing untuk mengelola alur logic aplikasi yang kompleks.

# Kekurangan dan Kelebihan

1. Kelebihan
   1. Single Responsibility Principle: Kita dapat mengekstrak komunikasi antar berbagai komponen ke dalam suatu tempat, sehingga lebih mudah untuk di maintain.
   2. Open/Closed Principle: Kita dapat memperkenalkan mediator baru tanpa harus mengubah komponen yang sebenarnya.
   3. Kita dapat mengurangin coupling antara berbagai komponen program.
   4. Kita dapat menggunakan kembari komponen individual dengan lebih mudah.
2. Kekurangan
   1. Seiring waktu, mediator dapat berevolusi menjadi God Object yang mengatur terlalu banyak hal.
