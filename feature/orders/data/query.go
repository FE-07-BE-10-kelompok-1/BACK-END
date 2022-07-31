package data

import (
	"bookstore/domain"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.OrderData {
	return &orderData{
		db: DB,
	}
}
