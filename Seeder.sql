START TRANSACTION;

INSERT INTO `Departemen` (`id_departemen`, `nama_departemen`) VALUES
(1, 'Teknologi Informasi'),
(2, 'Sumber Daya Manusia'),
(3, 'Pemasaran'),
(4, 'Keuangan');

INSERT INTO `Ruangan` (`id_ruangan`, `nama_ruangan`) VALUES
(1, 'Ruang Rapat Merapi'),
(2, 'Ruang Rapat Merbabu'),
(3, 'Auditorium Semeru');

INSERT INTO `Admin` (`id_admin`, `username`, `password`, `nama_lengkap`) VALUES
(1, 'admin_utama', 'hashed_password_1', 'Budi Santoso'),
(2, 'admin_cadangan', 'hashed_password_2', 'Citra Lestari');

INSERT INTO `Karyawan` (`id_karyawan`, `id_departemen`, `nama_karyawan`) VALUES
(1, 1, 'Andi Wijaya'),
(2, 1, 'Rina Permata'),
(3, 2, 'Dewi Anggraini'),
(4, 3, 'Eko Prasetyo'),
(5, 4, 'Sari Hartono');

INSERT INTO `Pemesanan` (`id_pemesanan`, `id_karyawan`, `id_ruangan`, `jam_dipinjam`, `jam_selesai`, `status`, `id_admin`) VALUES
(1, 1, 1, '2025-10-10 09:00:00', '2025-10-10 11:00:00', 'Menunggu Persetujuan', NULL),
(2, 4, 2, '2025-10-11 13:00:00', '2025-10-11 14:00:00', 'Menunggu Persetujuan', NULL),
(3, 2, 1, '2025-10-12 10:00:00', '2025-10-12 11:30:00', 'Disetujui', 1),
(4, 3, 3, '2025-10-12 14:00:00', '2025-10-12 17:00:00', 'Disetujui', 2),
(5, 5, 2, '2025-10-13 09:00:00', '2025-10-13 10:00:00', 'Ditolak', 1);

COMMIT;
