package usecase

import "bookstore/domain"

type bookUsecase struct {
	bookData domain.BookData
}

func New(bd domain.BookData) domain.BookUsecase {
	return &bookUsecase{
		bookData: bd,
	}
}
