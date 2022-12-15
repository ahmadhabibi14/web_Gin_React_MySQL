CREATE database crud_db

CREATE TABLE `books` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `release_year` year(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) 