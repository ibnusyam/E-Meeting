# E-Meeting API

Ini adalah API server untuk aplikasi E-Meeting, dibangun menggunakan Go (Golang) dengan framework Echo dan database PostgreSQL.

## 🚀 Fitur

-   Manajemen User
-   Manajemen Ruangan Rapat (Rooms)
-   Manajemen Snack
-   Sistem Reservasi
-   Dokumentasi API dengan Swagger

---

## ⚙️ Prasyarat

Sebelum memulai, pastikan Anda sudah menginstal perangkat lunak berikut:
-   [Go](https://golang.org/dl/) versi 1.20+
-   [PostgreSQL](https://www.postgresql.org/download/)
-   [migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
-   [Git](https://git-scm.com/)

---

## 🛠️ Langkah-langkah Instalasi

1.  **Clone Repositori**
    ```bash
    git clone [https://github.com/NAMA_USER_ANDA/NAMA_REPO_ANDA.git](https://github.com/NAMA_USER_ANDA/NAMA_REPO_ANDA.git)
    cd NAMA_REPO_ANDA
    ```

2.  **Siapkan Database**
    -   Buat sebuah database baru di PostgreSQL Anda (contoh: `e_meeting_db`).

3.  **Konfigurasi Environment**
    -   Salin file `.env.example` menjadi file baru bernama `.env`.
        ```bash
        cp .env.example .env
        ```
    -   Buka file `.env` dan sesuaikan nilainya, terutama untuk koneksi database (DB_USER, DB_PASSWORD, DB_NAME).

4.  **Install Dependensi**
    -   Jalankan perintah berikut untuk mengunduh semua library yang dibutuhkan.
        ```bash
        go mod tidy
        ```

5.  **Jalankan Migrasi Database**
    -   Perintah ini akan membuat semua tabel yang dibutuhkan di database Anda.
        ```bash
        # Pastikan Anda sudah menginstal migrate CLI
        migrate -database "postgres://USER:PASSWORD@HOST:PORT/DBNAME?sslmode=disable" -path migration up
        ```
    -   *Ganti `USER`, `PASSWORD`, `HOST`, `PORT`, `DBNAME` dengan konfigurasi Anda.*

6.  **Isi Data Awal (Seeding)**
    -   Jalankan perintah ini untuk mengisi database dengan data contoh.
        ```bash
        go run main.go --seed
        ```

---

## ▶️ Menjalankan Aplikasi

-   Untuk menjalankan server API, gunakan perintah:
    ```bash
    go run main.go
    ```
-   Server akan berjalan di `http://localhost:8080`.

## 📚 Dokumentasi API (Swagger)

-   Setelah server berjalan, dokumentasi API interaktif tersedia di:
    [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)