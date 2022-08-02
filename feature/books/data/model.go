package data

import (
	"bookstore/domain"
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

func ToEntity(data domain.Book) Book {
	return Book{
		Title:    data.Title,
		Image:    data.Image,
		Price:    data.Price,
		Stock:    data.Stock,
		Author:   data.Author,
		Sinopsis: data.Sinopsis,
		File:     data.File,
	}
}

func (b *Book) ToDomain() domain.Book {
	return domain.Book{
		ID:       b.ID,
		Title:    b.Title,
		Image:    b.Image,
		Price:    b.Price,
		Stock:    b.Stock,
		Author:   b.Author,
		Sinopsis: b.Sinopsis,
		File:     b.File,
	}
}
