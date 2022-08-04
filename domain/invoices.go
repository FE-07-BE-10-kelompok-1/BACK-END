package domain

import "time"

type Invoice struct {
	ID             string
	Users_ID       uint
	Total          int
	Status         string
	Payment_Link   string
	Payment_Method string
	Paid_At        string
}

type GetAllInvoices struct {
	ID             string               `json:"id" form:"id"`
	Username       string               `json:"username" form:"username"`
	Total          int                  `json:"total" form:"total"`
	Status         string               `json:"status" form:"status"`
	Payment_Link   string               `json:"payment_link" form:"payment_link"`
	Payment_Method string               `json:"payment_method" form:"payment_method"`
	Paid_At        string               `json:"paid_at" form:"paid_at"`
	Created_At     time.Time            `json:"created_at" form:"created_at"`
	Orders         []JoinOrderWithBooks `json:"orders" form:"orders"`
}

type InvoiceUsecase interface {
	CheckStocks([]uint) ([]Book, error)
	GetUserData(id uint) (User, error)
	UpdateStock([]uint) error
	InsertInvoice(data Invoice, id []uint) (Invoice, error)
	DeleteCarts(id uint) error
	GetAllOrders(user User) ([]GetAllInvoices, error)
	MidtransCallback(data Invoice) error
}

type InvoiceData interface {
	CheckStocks([]uint) ([]Book, error)
	GetUser(id uint) (User, error)
	UpdateStock([]uint) error
	Insert(data Invoice, id []uint) (Invoice, error)
	DeleteCarts(id uint) error
	GetAll() ([]GetAllInvoices, error)
	GetMyOrders(id uint) ([]GetAllInvoices, error)
	Update(data Invoice) error
}
