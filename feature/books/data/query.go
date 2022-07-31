package data

import (
	"bookstore/domain"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.BookData {
	return &bookData{
		db: DB,
	}
}
