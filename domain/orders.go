package domain

type Order struct {
	ID         uint
	Invoice_ID string
	Books_ID   uint
}

type JoinOrderWithBooks struct {
	ID         uint   `json:"id" form:"id"`
	Invoice_ID string `json:"invoice_id" form:"invoice_id"`
	Books_ID   uint   `json:"books_id" form:"books_id"`
	Title      string `json:"title" form:"title"`
	Image      string `json:"image" form:"image"`
	Price      uint   `json:"price" form:"price"`
}

type OrderUsecase interface{}

type OrderData interface{}
