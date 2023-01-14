-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: todo-mysql:3306
-- Generation Time: Jan 14, 2023 at 08:04 AM
-- Server version: 5.7.40
-- PHP Version: 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_todoapp`
--

INSERT INTO `permissions` (`id`, `name`, `display_name`, `description`) VALUES
('152220f1-93de-11ed-8ca1-0242ac1b0002', 'create_todo', 'Create Todo', 'Create Todo');


INSERT INTO `roles` (`id`, `name`, `display_name`, `description`) VALUES
('8a1af012-93de-11ed-8ca1-0242ac1b0002', 'ADMIN', 'admin', 'admin'),
('8a1b0c29-93de-11ed-8ca1-0242ac1b0002', 'USER', 'user', 'user');

INSERT INTO `users` (`username`, `password`, `is_active`) VALUES
('admin', '$2a$10$R8FDqmlE6/brL2SGPlPAXOUv3neiTRSB9osXzjO2sJmNh5tNe9jU2', 1),
('user', '$2a$10$R8FDqmlE6/brL2SGPlPAXOUv3neiTRSB9osXzjO2sJmNh5tNe9jU2', 1);

INSERT INTO `permission_role` (`permission_id`, `role_id`) VALUES
('152220f1-93de-11ed-8ca1-0242ac1b0002', '8a1af012-93de-11ed-8ca1-0242ac1b0002');


INSERT INTO `permission_user` (`permission_id`, `user_username`) VALUES
('152220f1-93de-11ed-8ca1-0242ac1b0002', 'admin');

INSERT INTO `role_user` (`role_id`, `user_username`) VALUES
('8a1af012-93de-11ed-8ca1-0242ac1b0002', 'admin'),
('8a1b0c29-93de-11ed-8ca1-0242ac1b0002', 'user');

COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
