-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 16, 2023 at 08:28 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tugasecho`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) NOT NULL,
  `name` varchar(255) NOT NULL,
  `age` int(10) NOT NULL,
  `address` varchar(255) NOT NULL,
  `country` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `age`, `address`, `country`) VALUES
(1, 'Joseph', 20, 'Jalan Jawa No 10', 'Indonesia'),
(2, 'Budi', 19, 'Jalan Bali No 11', 'Jepang'),
(3, 'Bimo', 21, 'Jalan Mangga No 19', 'Belanda'),
(4, 'Santi', 21, 'Jalan Mangga No 20', 'Belanda'),
(5, 'Yanti', 15, 'Jalan Sumedang No 110', 'Africa'),
(6, 'Vincent', 18, 'Jalan Kenari No 12', 'Indonesia'),
(7, 'Gerald', 18, 'Jalan Natuna No 13', 'Indonesia'),
(8, 'Opahh', 81, 'Jalan Sumatra No 24', 'Indonesia'),
(9, 'Ipin', 11, 'Jalan Sumatra No 24', 'Indonesia'),
(10, 'Opah', 80, 'Jalan Sumatra No 24', 'Indonesia');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
