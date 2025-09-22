-- phpMyAdmin SQL Dump
-- version 5.2.2deb1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Sep 22, 2025 at 11:50 AM
-- Server version: 8.4.6-0ubuntu0.25.04.3
-- PHP Version: 8.4.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `e-mee`
--

-- --------------------------------------------------------

--
-- Table structure for table `Admin`
--

CREATE TABLE `Admin` (
  `id_admin` int NOT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL COMMENT 'Password harus disimpan dalam bentuk hash',
  `nama_lengkap` varchar(150) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Departemen`
--

CREATE TABLE `Departemen` (
  `id_departemen` int NOT NULL,
  `nama_departemen` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Karyawan`
--

CREATE TABLE `Karyawan` (
  `id_karyawan` int NOT NULL,
  `id_departemen` int NOT NULL,
  `nama_karyawan` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Pemesanan`
--

CREATE TABLE `Pemesanan` (
  `id_pemesanan` int NOT NULL,
  `id_karyawan` int NOT NULL,
  `id_ruangan` int NOT NULL,
  `jam_dipinjam` datetime NOT NULL,
  `jam_selesai` datetime NOT NULL,
  `status` enum('Menunggu Persetujuan','Disetujui','Ditolak') NOT NULL DEFAULT 'Menunggu Persetujuan',
  `id_admin` int DEFAULT NULL COMMENT 'Admin yang menyetujui/menolak'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Ruangan`
--

CREATE TABLE `Ruangan` (
  `id_ruangan` int NOT NULL,
  `nama_ruangan` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Admin`
--
ALTER TABLE `Admin`
  ADD PRIMARY KEY (`id_admin`),
  ADD UNIQUE KEY `username` (`username`);

--
-- Indexes for table `Departemen`
--
ALTER TABLE `Departemen`
  ADD PRIMARY KEY (`id_departemen`),
  ADD UNIQUE KEY `nama_departemen` (`nama_departemen`);

--
-- Indexes for table `Karyawan`
--
ALTER TABLE `Karyawan`
  ADD PRIMARY KEY (`id_karyawan`),
  ADD KEY `id_departemen` (`id_departemen`);

--
-- Indexes for table `Pemesanan`
--
ALTER TABLE `Pemesanan`
  ADD PRIMARY KEY (`id_pemesanan`),
  ADD KEY `id_karyawan` (`id_karyawan`),
  ADD KEY `id_ruangan` (`id_ruangan`),
  ADD KEY `id_admin` (`id_admin`);

--
-- Indexes for table `Ruangan`
--
ALTER TABLE `Ruangan`
  ADD PRIMARY KEY (`id_ruangan`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `Admin`
--
ALTER TABLE `Admin`
  MODIFY `id_admin` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `Departemen`
--
ALTER TABLE `Departemen`
  MODIFY `id_departemen` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `Karyawan`
--
ALTER TABLE `Karyawan`
  MODIFY `id_karyawan` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `Pemesanan`
--
ALTER TABLE `Pemesanan`
  MODIFY `id_pemesanan` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `Ruangan`
--
ALTER TABLE `Ruangan`
  MODIFY `id_ruangan` int NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `Karyawan`
--
ALTER TABLE `Karyawan`
  ADD CONSTRAINT `karyawan_ibfk_1` FOREIGN KEY (`id_departemen`) REFERENCES `Departemen` (`id_departemen`);

--
-- Constraints for table `Pemesanan`
--
ALTER TABLE `Pemesanan`
  ADD CONSTRAINT `pemesanan_ibfk_1` FOREIGN KEY (`id_karyawan`) REFERENCES `Karyawan` (`id_karyawan`),
  ADD CONSTRAINT `pemesanan_ibfk_2` FOREIGN KEY (`id_ruangan`) REFERENCES `Ruangan` (`id_ruangan`),
  ADD CONSTRAINT `pemesanan_ibfk_3` FOREIGN KEY (`id_admin`) REFERENCES `Admin` (`id_admin`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
