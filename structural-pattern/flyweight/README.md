# Flyweight Pattern

Flyweight pattern adalah sebuah structural design pattern yang memungkinkan program untuk support object dalam jumlah besar dengan menjaga memory tetap rendah
Flyweight pattern ini mencapainya dengan berbagi bagian dari status obect di antara beberapa object. Dengan kata lain, Flyweight menghemat RAM dengan menyimpan data yang sama yang digunakan oleh object yang berbeda.

## Conceptual Example

Dalam sebuah game counter strike, Terrorist dan Counter-Terrorist memiliki sebuah perpedaan type pakaian. Untuk simple nya.
Asumsikan bahwa kedua Terrorist dan Counter-Terrorist memiliki type pakaian. Object pakaian terembedded(tertanam) dalam sebuah object pemain

```go
type Player struct {
    dress dress
    playerType string // can be T or CT
    lat int
    long int
}
```

Lets say kita memiliki 5 Terrorist dan 5 Counter-Terrorist, jadi total 10 pemain. Sekarang ada 2 pilihan mengenai pakaian.

1. Masing2 dari 10 object pemain menciptakan object pakaian yang berbeda dan menyematkannya. Sebanyak 10 object pakaian yang akan di buat.
2. Kita membuat 2 object pakaian:
   - Single Terroristt object pakaian. Ini akan dibagikan kepada 5 terrorist
   - Single Counter-Terrorist object pakaian. Ini akan di bagikan kepada 5 counter-terrorist

Seperti yang kamu lihat pada pendekatan 1, total ada 10 object pakaian yang dibuat, sedangkan pada pendekatan 2 hanya 2 object pakaian yang dibuat. Pendekatan ke2 adalah apa yang kita ikut dalam design pattern flyweight. 2 Object pakaian yang kami buat di sebut object flyweight.

Pattern flyweight menghilangkan bagian-bagian umum dan menciptakan object2 flyweight. Object flyweight ini (pakaian) kemudian dapat dibagi di antara beberapa object pemain. Hal ini secara drastis mengurangi jumlah pada object pakaian, dan bagian baiknya adalah meskipun kmamu membuat lebih banya pemain, hanya 2 object pakain saja yang cukup.

Pada pattern flyweight, kita menyimpan object flyweight pada bidang peta. Setiap kali object lain yang memiliki object flyweight dibuat, maka object flyweight akan diambil dari peta.

Mari kita liat, bagian mana dari pengaturan ini yang akan menjadi status intrinsik dan ekstrinsik.

1. Intrinsik state: Berpakaian dalam kondisi intrinsik karena dapat digunakan dibeberapa object Terrorist dan Counter-Terrorist.
2. Ekstrinsik state: Lokasi pemain dan senjata pemain adalah status ekstrinsik karena berbeda untuk setiap object.
