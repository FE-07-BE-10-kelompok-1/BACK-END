package data

import (
	cartData "bookstore/feature/carts/data"
	"bookstore/feature/invoices/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Username string
	Phone    string
	Password string
	Role     string
	Carts    []cartData.Cart `gorm:"foreignKey:Users_ID"`
	Invoices []data.Invoice  `gorm:"foreignKey:Users_ID"`
}
