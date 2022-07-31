package data

import (
	"bookstore/domain"

	"gorm.io/gorm"
)

type invoiceData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.InvoiceData {
	return &invoiceData{
		db: DB,
	}
}
