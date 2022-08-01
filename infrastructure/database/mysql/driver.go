package mysql

import (
	"bookstore/config"
	booksData "bookstore/feature/books/data"
	cartsData "bookstore/feature/carts/data"
	invoicesData "bookstore/feature/invoices/data"
	ordersData "bookstore/feature/orders/data"
	usersData "bookstore/feature/users/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(usersData.User{}, booksData.Book{}, cartsData.Cart{}, ordersData.Order{}, invoicesData.Invoice{})
}
