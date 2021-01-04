CREATE DATABASE IF NOT EXISTS `book-shelf-test`;
USE `book-shelf-test`;

CREATE TABLE IF NOT EXISTS users (
	`id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`email` VARCHAR(255) NOT NULL,
	`password_hash` VARCHAR(255) NOT NULL,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS user_session (
    `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `token` VARBINARY(512) NOT NULL,
    `expires_at` DATETIME(6) NOT NULL,
    `created_at` DATETIME(6) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS books (
	`id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(255) NOT NULL,
	`author` VARCHAR(255) NOT NULL,
    `url` VARCHAR(255) NOT NULL,
    `release_date` TIMESTAMP NULL,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS users_books (
	`id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	`user_id` BIGINT NOT NULL,
	`book_id` BIGINT NOT NULL,
    `finish_date` TIMESTAMP NOT NULL,
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,
    FOREIGN KEY (user_id) 
        REFERENCES users(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (book_id) 
        REFERENCES books(id)
        ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;
