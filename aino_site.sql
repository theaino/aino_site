-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 20, 2023 at 08:49 PM
-- Server version: 11.2.2-MariaDB
-- PHP Version: 8.2.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `aino_site`
--

-- --------------------------------------------------------

--
-- Table structure for table `pages`
--

CREATE TABLE `pages` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` longtext DEFAULT NULL,
  `template` longtext DEFAULT NULL,
  `is_admin_page` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `pages`
--

INSERT INTO `pages` (`id`, `created_at`, `updated_at`, `deleted_at`, `path`, `template`, `is_admin_page`) VALUES
(1, NULL, NULL, NULL, '/about', 'about', 0),
(2, NULL, NULL, NULL, '/', 'home', 0),
(3, NULL, NULL, NULL, '/home', 'home', 0),
(4, NULL, NULL, NULL, '/', 'admin-home', 1),
(5, NULL, NULL, NULL, '/login', 'login', 0),
(6, NULL, NULL, NULL, '/contact', 'contact', 0);

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `contents` longtext DEFAULT NULL,
  `date` datetime(3) DEFAULT NULL,
  `public` tinyint(1) DEFAULT NULL,
  `abstract` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `contents`, `date`, `public`, `abstract`) VALUES
(1, '0000-00-00 00:00:00.000', '2023-12-20 21:49:01.847', NULL, 'Hello There!', 'This is a test.<h1>Hello!</h1>General Kenobi!', '2023-12-17 16:39:21.000', 1, 'The first post which is a test.Im aino'),
(2, NULL, NULL, NULL, '2nd', 'This is a private test.', '2023-12-17 16:40:19.000', 0, 'Test post but it is private.'),
(3, NULL, NULL, NULL, 'Over the Rainbow', '<button>Le Button</button>\r\n\r\n<marquee>Guten Tag</marquee>', NULL, 1, 'Copyrighyy');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `pages`
--
ALTER TABLE `pages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_pages_deleted_at` (`deleted_at`);

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_posts_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `pages`
--
ALTER TABLE `pages`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
