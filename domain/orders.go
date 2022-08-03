package domain

type Order struct {
	ID         uint
	Invoice_ID string
	Books_ID   uint
}

type OrderUsecase interface{}

type OrderData interface{}
