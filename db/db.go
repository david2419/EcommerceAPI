package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

// use ecommerce_db;
// CREATE TABLE IF NOT EXISTS users (
//   `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
//   `firstName` VARCHAR(255) NOT NULL,
//   `lastName` VARCHAR(255) NOT NULL,
//   `email` VARCHAR(255) NOT NULL,
//   `password` VARCHAR(255) NOT NULL,
//   `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

//   PRIMARY KEY (id),
//   UNIQUE KEY (email)
// );
// CREATE TABLE IF NOT EXISTS products (
//   `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
//   `name` VARCHAR(255) NOT NULL,
//   `description` TEXT NOT NULL,
//   `image` VARCHAR(255) NOT NULL,
//   `price` DECIMAL(10, 2) NOT NULL,
//   `quantity` INT UNSIGNED NOT NULL,
//   `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   PRIMARY KEY (id)
// );
// CREATE TABLE IF NOT EXISTS orders (
//   `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
//   `userId` INT UNSIGNED NOT NULL,
//   `total` DECIMAL(10, 2) NOT NULL,
//   `status` ENUM('pending', 'completed', 'cancelled') NOT NULL DEFAULT 'pending',
//   `address` TEXT NOT NULL,
//   `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

//   PRIMARY KEY (`id`),
//   FOREIGN KEY (`userId`) REFERENCES users(`id`)
// );
// CREATE TABLE IF NOT EXISTS `order_items` (
//   `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
//   `orderId` INT UNSIGNED NOT NULL,
//   `productId` INT UNSIGNED NOT NULL,
//   `quantity` INT NOT NULL,
//   `price` DECIMAL(10, 2) NOT NULL,

//   PRIMARY KEY (`id`),
//   FOREIGN KEY (`orderId`) REFERENCES orders(`id`),
//   FOREIGN KEY (`productId`) REFERENCES products(`id`)
// );
