package usecase

import "bookstore/domain"

type cartUsecase struct {
	cartData domain.CartData
}

func New(cd domain.CartData) domain.CartUsecase {
	return &cartUsecase{
		cartData: cd,
	}
}
