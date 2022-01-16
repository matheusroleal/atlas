DROP DATABASE IF EXISTS `Atlas`;

CREATE DATABASE `Atlas`;

USE `Atlas`;

DROP TABLE IF EXISTS `Segments`;

CREATE TABLE `Segments` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `Owner` varchar(256) DEFAULT NULL,
  `Data` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `Tracks`;

CREATE TABLE `Tracks` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `Owner` varchar(256) DEFAULT NULL,
  `Data` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;