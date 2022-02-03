-- create database if not exist
CREATE DATABASE IF NOT EXISTS `eventkoe-project`;

-- create table users
DROP TABLE IF EXISTS `eventkoe-project`.`users`;
CREATE TABLE `eventkoe-project`.`users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(1000) NOT NULL,
  `birth_date` date NOT NULL,
  `phone_number` varchar(50) NOT NULL,
  `photo` varchar(1000),
  `gender` varchar(10) NOT NULL,
  `address` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT NULL,
  `updated_at` timestamp DEFAULT NULL,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table category
DROP TABLE IF EXISTS `eventkoe-project`.`category`;
CREATE TABLE `eventkoe-project`.`category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `description` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT NULL,
  `updated_at` timestamp DEFAULT NULL,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY(`id`)
);

-- create table event
DROP TABLE IF EXISTS `eventkoe-project`.`event`;
CREATE TABLE `eventkoe-project`.`event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_user` int NOT NULL,
  `id_category` int NOT NULL,
  `description` varchar(255) NOT NULL,
  `start_date` timestamp NOT NULL,
  `end_date` timestamp NOT NULL,
  `location` varchar(255) NOT NULL,
  `details` varchar(255) NOT NULL,
  `photo` varchar(1000),
  `created_at` timestamp DEFAULT NULL,
  `updated_at` timestamp DEFAULT NULL,
  `deleted_at` timestamp DEFAULT NULL,
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
  `created_at` timestamp DEFAULT NULL,
  `updated_at` timestamp DEFAULT NULL,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`id_user`) REFERENCES users(`id`),
  FOREIGN KEY(`id_event`) REFERENCES event(`id`)
);

-- create table participant
DROP TABLE IF EXISTS `eventkoe-project`.`participant`;
CREATE TABLE `eventkoe-project`.`participant` (
  `id` int NOT NULL AUTO_INCREMENT,
  `id_event` int NOT NULL,
  `id_user` int NOT NULL,
  `created_at` timestamp DEFAULT NULL,
  `updated_at` timestamp DEFAULT NULL,
  `deleted_at` timestamp DEFAULT NULL,
  PRIMARY KEY(`id`),
  FOREIGN KEY(`id_event`) REFERENCES event(`id`),
  FOREIGN KEY(id_user) REFERENCES users(`id`)
);