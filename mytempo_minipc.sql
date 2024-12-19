-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Dec 19, 2024 at 07:24 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `mytempo_minipc`
--

-- --------------------------------------------------------

--
-- Table structure for table `athletes`
--

CREATE TABLE IF NOT EXISTS `athletes` (
  `num` int(11) NOT NULL,
  `event_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `team` varchar(80) DEFAULT NULL,
  `track_id` int(11) DEFAULT NULL,
  `sex` varchar(1) NOT NULL,
  PRIMARY KEY (`num`),
  KEY `event_id` (`event_id`),
  KEY `track_id` (`track_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `athletes_times`
--

CREATE TABLE IF NOT EXISTS `athletes_times` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `antenna` int(11) DEFAULT NULL,
  `checkpoint_id` int(11) NOT NULL,
  `athlete_num` int(11) NOT NULL,
  `athlete_time` varchar(12) DEFAULT NULL,
  `staff` int(11) NOT NULL,
  `timestp` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `athlete_num` (`athlete_num`),
  KEY `athlete_time` (`athlete_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `checkpoints`
--

CREATE TABLE IF NOT EXISTS `checkpoints` (
  `id` int(11) NOT NULL,
  `description` int(11) NOT NULL,
  `km` int(11) NOT NULL,
  `local` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `classificacao`
--

CREATE TABLE IF NOT EXISTS `classificacao` (
  `athlete` int(11) DEFAULT NULL,
  `start_time` varchar(12) DEFAULT NULL,
  `end_time` varchar(12) DEFAULT NULL,
  UNIQUE KEY `athlete_2` (`athlete`),
  KEY `athlete` (`athlete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `equipamento`
--

CREATE TABLE IF NOT EXISTS `equipamento` (
  `id` tinyint(1) NOT NULL,
  `checkpoint_id` int(11) NOT NULL,
  `idequip` int(11) NOT NULL,
  `modelo` varchar(30) NOT NULL,
  `event_id` int(11) NOT NULL,
  `tags_unicas` int(11) NOT NULL DEFAULT 0,
  `action` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `event_id` (`event_id`),
  KEY `checkpoint_id` (`checkpoint_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `event_data`
--

CREATE TABLE IF NOT EXISTS `event_data` (
  `id` int(11) NOT NULL,
  `event_date` date DEFAULT NULL,
  `event_title` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `invalidos`
--

CREATE TABLE IF NOT EXISTS `invalidos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `antenna` int(11) DEFAULT NULL,
  `checkpoint_id` int(11) NOT NULL,
  `athlete_num` int(11) NOT NULL,
  `athlete_time` varchar(12) DEFAULT NULL,
  `staff` int(11) NOT NULL,
  `timestp` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `athlete_num` (`athlete_num`),
  KEY `athlete_time` (`athlete_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `rede`
--

CREATE TABLE IF NOT EXISTS `rede` (
  `ssid` varchar(100) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(255) DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT 0,
  `descricao` varchar(255) NOT NULL DEFAULT 'no description',
  `conectar` tinyint(4) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `staffs`
--

CREATE TABLE IF NOT EXISTS `staffs` (
  `id` int(11) NOT NULL,
  `event_id` int(11) NOT NULL,
  `nome` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `event_id` (`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `tracks`
--

CREATE TABLE IF NOT EXISTS `tracks` (
  `id` int(11) NOT NULL,
  `event_id` int(11) DEFAULT NULL,
  `race_description` varchar(255) DEFAULT NULL,
  `inicio` time DEFAULT NULL,
  `chegada` time DEFAULT NULL,
  `largada` time DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `event_id` (`event_id`),
  KEY `inicio` (`inicio`),
  KEY `chegada` (`chegada`),
  KEY `largada` (`largada`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `unsent_athletes`
--

CREATE TABLE IF NOT EXISTS `unsent_athletes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `antenna` int(11) DEFAULT NULL,
  `checkpoint_id` int(11) NOT NULL,
  `athlete_num` int(11) NOT NULL,
  `athlete_time` varchar(12) DEFAULT NULL,
  `staff` int(11) NOT NULL,
  `timestp` timestamp NOT NULL DEFAULT current_timestamp(),
  `percurso` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `athlete_num` (`athlete_num`),
  KEY `percursoid` (`percurso`),
  KEY `percurso` (`percurso`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `athletes`
--
ALTER TABLE `athletes`
  ADD CONSTRAINT `athletes_ibfk_1` FOREIGN KEY (`event_id`) REFERENCES `event_data` (`id`);

--
-- Constraints for table `athletes_times`
--
ALTER TABLE `athletes_times`
  ADD CONSTRAINT `athletes_times_ibfk_1` FOREIGN KEY (`athlete_num`) REFERENCES `athletes` (`num`);

--
-- Constraints for table `classificacao`
--
ALTER TABLE `classificacao`
  ADD CONSTRAINT `classificacao_ibfk_1` FOREIGN KEY (`athlete`) REFERENCES `athletes_times` (`athlete_num`);

--
-- Constraints for table `staffs`
--
ALTER TABLE `staffs`
  ADD CONSTRAINT `staffs_ibfk_1` FOREIGN KEY (`event_id`) REFERENCES `event_data` (`id`);

--
-- Constraints for table `tracks`
--
ALTER TABLE `tracks`
  ADD CONSTRAINT `tracks_ibfk_1` FOREIGN KEY (`event_id`) REFERENCES `event_data` (`id`);

--
-- Constraints for table `unsent_athletes`
--
ALTER TABLE `unsent_athletes`
  ADD CONSTRAINT `unsent_athletes_ibfk_1` FOREIGN KEY (`athlete_num`) REFERENCES `athletes` (`num`),
  ADD CONSTRAINT `unsent_athletes_ibfk_2` FOREIGN KEY (`percurso`) REFERENCES `tracks` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
