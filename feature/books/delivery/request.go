package delivery

import "bookstore/domain"

type BookRequest struct {
	Title    string `json:"title" form:"title" validate:"required"`
	Image    string `json:"image" form:"image"`
	Price    int    `json:"price" form:"price" validate:"required"`
	Stock    int    `json:"stock" form:"stock" validate:"required"`
	Author   string `json:"author" form:"author" validate:"required"`
	Sinopsis string `json:"sinopsis" form:"sinopsis" validate:"required"`
	File     string `json:"file" form:"file"`
}

func (br *BookRequest) ToDomain() domain.Book {
	return domain.Book{
		Title:    br.Title,
		Image:    br.Image,
		Price:    br.Price,
		Stock:    br.Stock,
		Author:   br.Author,
		Sinopsis: br.Sinopsis,
		File:     br.File,
	}
}

type BookUpdateRequest struct {
	Title    string `json:"title" form:"title"`
	Price    int    `json:"price" form:"price"`
	Stock    int    `json:"stock" form:"stock"`
	Author   string `json:"author" form:"author"`
	Sinopsis string `json:"sinopsis" form:"sinopsis"`
}

func (br *BookUpdateRequest) ToDomain() domain.Book {
	return domain.Book{
		Title:    br.Title,
		Price:    br.Price,
		Stock:    br.Stock,
		Author:   br.Author,
		Sinopsis: br.Sinopsis,
	}
}
