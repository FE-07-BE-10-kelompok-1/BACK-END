package domain

type Cart struct {
	ID       uint
	Books_ID uint
	Users_ID uint
}

type CartUsecase interface{}

type CartData interface{}
