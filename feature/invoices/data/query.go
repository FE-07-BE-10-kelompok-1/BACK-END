package data

import (
	"bookstore/domain"
	"errors"

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

func (id *invoiceData) CheckStocks(booksID []uint) ([]domain.Book, error) {
	var books []domain.Book
	id.db.Where("id in (?)", booksID).Find(&books)
	if len(books) != len(booksID) {
		return []domain.Book{}, errors.New("books result not matched, try refreshing the carts page")
	}
	for i := 0; i < len(books); i++ {
		if books[i].Stock-1 < 0 {
			return []domain.Book{}, errors.New("book stock of " + books[i].Title + " is 0, delete book from your cart")
		}
	}

	return books, nil
}

func (id *invoiceData) GetUser(userID uint) (domain.User, error) {
	var userData domain.User
	err := id.db.Where("id = ?", userID).First(&userData).Error
	if err != nil {
		return domain.User{}, err
	}

	return userData, nil
}

func (id *invoiceData) UpdateStock(booksID []uint) error {
	err := id.db.Exec("UPDATE books SET stock = stock - 1, updated_at = now() WHERE id in (?)", booksID).Error
	return err
}

func (id *invoiceData) Insert(data domain.Invoice, booksID []uint) (domain.Invoice, error) {
	var invoiceData Invoice = ToEntity(data)
	err := id.db.Create(&invoiceData).Error
	if err != nil {
		return domain.Invoice{}, err
	}

	for i := 0; i < len(booksID); i++ {
		var order = domain.Order{Invoice_ID: invoiceData.ID, Books_ID: booksID[i]}
		err := id.db.Create(&order).Error
		if err != nil {
			return domain.Invoice{}, err
		}
	}

	return invoiceData.ToDomain(), nil
}

func (id *invoiceData) DeleteCarts(userID uint) error {
	err := id.db.Exec("UPDATE carts SET deleted_at = now() WHERE users_id = ?", userID).Error
	return err
}
