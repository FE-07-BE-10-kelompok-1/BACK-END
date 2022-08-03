package domain

import "time"

type Invoice struct {
	ID             string
	Users_ID       uint
	Total          int
	Status         string
	Payment_Link   string
	Payment_Method string
	Paid_At        time.Time
}

type InvoiceUsecase interface {
	CheckStocks([]uint) ([]Book, error)
	GetUserData(id uint) (User, error)
	UpdateStock([]uint) error
	InsertInvoice(data Invoice, id []uint) (Invoice, error)
}

type InvoiceData interface {
	CheckStocks([]uint) ([]Book, error)
	GetUser(id uint) (User, error)
	UpdateStock([]uint) error
	Insert(data Invoice, id []uint) (Invoice, error)
}
