package domain

type Book struct {
	ID       uint
	Title    string
	Image    string
	Price    int
	Stock    int
	Author   string
	Sinopsis string
	File     string
}

type BookUsecase interface{}

type BookData interface{}
