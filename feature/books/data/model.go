package data

import (
	cartData "bookstore/feature/carts/data"
	"bookstore/feature/orders/data"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title    string
	Image    string
	Price    int
	Stock    int
	Author   string
	Sinopsis string
	File     string
	Carts    []cartData.Cart `gorm:"foreignKey:Books_ID"`
	Order    []data.Order    `gorm:"foreignKey:Books_ID"`
}
