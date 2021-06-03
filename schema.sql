CREATE TABLE `user` (
  `id` bigint PRIMARY KEY,
  `nik` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `jenisKelamin` varchar(255) NOT NULL,
  `tanggalLahir` varchar(255) NOT NULL,
  `agama` varchar(255) NOT NULL,
  `createdAt` timestamptz NOT NULL DEFAULT (now())
  `updatedAt` timestamptz 
);
