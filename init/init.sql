-- create database if not exist
CREATE DATABASE IF NOT EXISTS `eventkoe-project`;

-- create table users
DROP TABLE IF EXISTS `eventkoe-project`.`users`;
CREATE TABLE `eventkoe-project`.`users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(1000) NOT NULL,
  `birth_date` varchar(255),
  `phone_number` varchar(50),
  `photo` varchar(1000),
  `gender` varchar(10),
  `address` varchar(255),
  `created_at` DATETIME DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table category
DROP TABLE IF EXISTS `eventkoe-project`.`category`;
CREATE TABLE `eventkoe-project`.`category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `created_at` DATETIME DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table event
DROP TABLE IF EXISTS `eventkoe-project`.`event`;
CREATE TABLE `eventkoe-project`.`event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_user` int NOT NULL,
  `id_category` int NOT NULL,
  `title` varchar(255) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `location` varchar(255) NOT NULL,
  `details` varchar(255) NOT NULL,
  `photo` varchar(1000),
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`id_user`) REFERENCES users(`id`),
  FOREIGN KEY(`id_category`) REFERENCES category(`id`)
);

-- create table comment
DROP TABLE IF EXISTS `eventkoe-project`.`comment`;
CREATE TABLE `eventkoe-project`.`comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_event` int NOT NULL,
  `id_user` int NOT NULL,
  `comment` varchar(255),
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`id_user`) REFERENCES users(`id`),
  FOREIGN KEY(`id_event`) REFERENCES `event`(`id`)
);

-- create table participant
DROP TABLE IF EXISTS `eventkoe-project`.`participant`;
CREATE TABLE `eventkoe-project`.`participant` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_event` int NOT NULL,
  `id_user` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`id_event`) REFERENCES `event`(`id`),
  FOREIGN KEY(`id_user`) REFERENCES users(`id`)
);