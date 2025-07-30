CREATE DATABASE IF NOT EXISTS `test_db`;

USE `test_db`;

CREATE TABLE `users` (
 `id` int NOT NULL AUTO_INCREMENT,
 `name` varchar(255) DEFAULT NULL,
 `address` varchar(255) DEFAULT NULL,
 `country` varchar(255) DEFAULT NULL,
 PRIMARY KEY (`id`)
);