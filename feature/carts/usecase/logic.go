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

func (cs *cartUsecase) AddToCart(data domain.Cart) error {
	err := cs.cartData.Insert(data)
	return err
}

func (cs *cartUsecase) GetCarts(id uint) ([]domain.JoinCartWithBooks, error) {
	data, err := cs.cartData.GetAll(id)
	return data, err
}

func (cs *cartUsecase) DeleteFromCart(data domain.Cart) error {
	err := cs.cartData.Delete(data)
	return err
}
