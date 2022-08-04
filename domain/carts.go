package domain

type Cart struct {
	ID       uint
	Books_ID uint
	Users_ID uint
}

type JoinCartWithBooks struct {
	ID       uint
	Books_ID uint
	Title    string
	Price    uint
	Image    string
	Author   string
}

type CartUsecase interface {
	AddToCart(data Cart) error
	GetCarts(id uint) ([]JoinCartWithBooks, error)
	DeleteFromCart(data Cart) error
}

type CartData interface {
	Insert(data Cart) error
	GetAll(id uint) ([]JoinCartWithBooks, error)
	Delete(data Cart) error
}
