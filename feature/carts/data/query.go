package data

import (
	"bookstore/domain"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.CartData {
	return &cartData{
		db: DB,
	}
}
