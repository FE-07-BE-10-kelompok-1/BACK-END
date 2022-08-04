package usecase

import (
	"bookstore/domain"
	"bookstore/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddToCart(t *testing.T) {
	repo := mocks.CartData{}
	usecase := New(&repo)
	successAdd := domain.Cart{Books_ID: 1, Users_ID: 3}

	t.Run("success add to cart", func(t *testing.T) {
		repo.On("Insert", successAdd).Return(nil).Once()
		err := usecase.AddToCart(successAdd)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetCarts(t *testing.T) {
	repo := mocks.CartData{}
	usecase := New(&repo)
	successGet := []domain.JoinCartWithBooks{{ID: 1, Books_ID: 1, Title: "yohohoh", Price: 20000, Image: "url.com", Author: "Oda sensei"}}

	t.Run("success get carts", func(t *testing.T) {
		repo.On("GetAll", uint(1)).Return(successGet, nil).Once()
		data, err := usecase.GetCarts(uint(1))
		assert.Nil(t, err)
		assert.Greater(t, len(data), 0)
		repo.AssertExpectations(t)
	})
}

func TestDeleteFromCart(t *testing.T) {
	repo := mocks.CartData{}
	usecase := New(&repo)
	deleteReq := domain.Cart{ID: 1, Books_ID: 1, Users_ID: 3}

	t.Run("success delete from cart", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()
		err := usecase.DeleteFromCart(deleteReq)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
