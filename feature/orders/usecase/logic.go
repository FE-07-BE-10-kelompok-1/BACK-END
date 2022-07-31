package usecase

import "bookstore/domain"

type orderUsecase struct {
	orderData domain.OrderData
}

func New(od domain.OrderData) domain.OrderUsecase {
	return &orderUsecase{
		orderData: od,
	}
}
