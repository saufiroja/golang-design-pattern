# Proxy Pattern

Proxy Pattern adalah sebuah structural design pattern yang menyediakan sebuah object yang bertindak sebagai pengganti object real service yang digunakan oleh client. Proxy menerima permintaan client, melakukan beberapa task (control access, caching, dll) dan kemudian meneruskan permintaan tersevut ke object real service.

## Conceptual Example

Web server seperti Nginx dapat bertindak sebagai proxy untuk aplikasi:

- Menyediakan control access ke aplikasi
- Dapat melakukan rate limiting
- Dapat melakukan caching
