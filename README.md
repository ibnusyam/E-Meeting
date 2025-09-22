# Skema Database Aplikasi Pemesanan Ruangan

Selamat datang di repositori skema database untuk aplikasi pemesanan ruangan. Proyek ini berisi file-file SQL yang dibutuhkan untuk membuat dan mengisi struktur database untuk sistem manajemen pemesanan ruangan kantor.

## Struktur Tabel

Database ini terdiri dari lima tabel utama yang saling berhubungan untuk mengelola data departemen, karyawan, ruangan, admin, dan transaksi pemesanan.

### 🏢 `Departemen`
Tabel ini berfungsi sebagai master data untuk menyimpan semua departemen yang ada di perusahaan.
- **`id_departemen`**: Kunci utama (Primary Key) untuk setiap departemen.
- **`nama_departemen`**: Nama unik untuk setiap departemen (contoh: 'Teknologi Informasi', 'Sumber Daya Manusia').

### 👨‍💼 `Karyawan`
Tabel ini menyimpan data semua karyawan yang dapat melakukan pemesanan ruangan.
- **`id_karyawan`**: Kunci utama (Primary Key) untuk setiap karyawan.
- **`id_departemen`**: Kunci asing (Foreign Key) yang menghubungkan karyawan ke tabel `Departemen`.
- **`nama_karyawan`**: Nama lengkap karyawan.

### 🚪 `Ruangan`
Tabel ini adalah master data untuk semua ruangan yang bisa dipesan.
- **`id_ruangan`**: Kunci utama (Primary Key) untuk setiap ruangan.
- **`nama_ruangan`**: Nama ruangan (contoh: 'Ruang Rapat Merapi').

### 🔑 `Admin`
Tabel ini berisi data pengguna dengan hak akses admin yang dapat menyetujui atau menolak pemesanan.
- **`id_admin`**: Kunci utama (Primary Key) untuk setiap admin.
- **`username`**: Username unik untuk login admin.
- **`password`**: Password admin (sebaiknya disimpan dalam bentuk hash).
- **`nama_lengkap`**: Nama lengkap admin.

### 📝 `Pemesanan`
Ini adalah tabel transaksi utama yang mencatat semua aktivitas pemesanan ruangan.
- **`id_pemesanan`**: Kunci utama (Primary Key) untuk setiap transaksi pemesanan.
- **`id_karyawan`**: Kunci asing (Foreign Key) yang menunjukkan siapa yang memesan.
- **`id_ruangan`**: Kunci asing (Foreign Key) yang menunjukkan ruangan mana yang dipesan.
- **`jam_dipinjam`** & **`jam_selesai`**: Waktu mulai dan selesai pemesanan.
- **`status`**: Status pemesanan (`Menunggu Persetujuan`, `Disetujui`, `Ditolak`).
- **`id_admin`**: Kunci asing (Foreign Key) yang mencatat admin mana yang memproses pemesanan.

---

## Visualisasi ERD (Entity-Relationship Diagram)

Untuk melihat visualisasi hubungan antar tabel, Anda dapat mengunjungi tautan dbdiagram di bawah ini:

[**Lihat Diagram Database di dbdiagram.io**](https://dbdiagram.io/d/68d136f07c85fb9961cd2049)
