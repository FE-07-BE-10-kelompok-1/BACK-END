package usecase

import (
	"bookstore/domain"
	"bookstore/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckStocks(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)
	successData := []domain.Book{{ID: 1, Title: "One Piece", Image: "url.com", Price: 20000, Stock: 1, Author: "Oda", Sinopsis: "Kaizoku", File: "url.com"}}

	t.Run("success case", func(t *testing.T) {
		repo.On("CheckStocks", []uint{1}).Return(successData, nil).Once()
		data, err := usecase.CheckStocks([]uint{1})
		assert.Nil(t, err)
		assert.Greater(t, len(data), 0)
		repo.AssertExpectations(t)
	})
}

func TestGetUserData(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)
	successCase := domain.User{ID: 1, Fullname: "athadf", Username: "atha", Phone: "089777", Password: "asdasda", Role: "admin"}

	t.Run("success case", func(t *testing.T) {
		repo.On("GetUser", uint(1)).Return(successCase, nil).Once()
		data, err := usecase.GetUserData(uint(1))
		assert.Nil(t, err)
		assert.Greater(t, data.ID, uint(0))
		repo.AssertExpectations(t)
	})
}

func TestUpdateStock(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)

	t.Run("success case", func(t *testing.T) {
		repo.On("UpdateStock", []uint{1}).Return(nil).Once()
		err := usecase.UpdateStock([]uint{1})
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestInsertInvoice(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)
	successCase := domain.Invoice{ID: "sdsfsdfsdf", Users_ID: 2, Total: 40000, Status: "waiting", Payment_Link: "url.com", Payment_Method: "gopay", Paid_At: ""}

	t.Run("success case", func(t *testing.T) {
		repo.On("Insert", mock.Anything, []uint{1}).Return(successCase, nil).Once()
		data, err := usecase.InsertInvoice(successCase, []uint{1})
		assert.Nil(t, err)
		assert.Greater(t, data.Users_ID, uint(0))
		repo.AssertExpectations(t)
	})
}

func TestDeleteCarts(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)

	t.Run("success case", func(t *testing.T) {
		repo.On("DeleteCarts", uint(1)).Return(nil).Once()
		err := usecase.DeleteCarts(uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAllOrders(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)
	successData := []domain.GetAllInvoices{}

	t.Run("success case admin", func(t *testing.T) {
		repo.On("GetAll").Return(successData, nil).Once()
		data, err := usecase.GetAllOrders(domain.User{ID: 1, Role: "admin"})
		assert.Nil(t, err)
		assert.NotNil(t, data)
		repo.AssertExpectations(t)
	})

	t.Run("success case user", func(t *testing.T) {
		repo.On("GetMyOrders", uint(1)).Return(successData, nil).Once()
		data, err := usecase.GetAllOrders(domain.User{ID: 1, Role: "user"})
		assert.Nil(t, err)
		assert.NotNil(t, data)
		repo.AssertExpectations(t)
	})
}

func TestMidtransCallback(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)
	mockData := domain.Invoice{ID: "sadasda", Users_ID: 1, Total: 20000, Status: "waiting", Payment_Link: "url.com", Payment_Method: "bca", Paid_At: ""}

	t.Run("success case", func(t *testing.T) {
		repo.On("Update", mockData, "asdasdas").Return(nil).Once()
		err := usecase.MidtransCallback(mockData, "asdasdas")
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetOrder(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)

	t.Run("success case", func(t *testing.T) {
		repo.On("GetOrder", "sadasdas", uint(1)).Return(nil).Once()
		err := usecase.GetOrder("sadasdas", uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdateStockAfterCancel(t *testing.T) {
	repo := mocks.InvoiceData{}
	usecase := New(&repo)

	t.Run("success case", func(t *testing.T) {
		repo.On("UpdateStockAfterCancel", "yahahaha").Return(nil).Once()
		err := usecase.UpdateStockAfterCancel("yahahaha")
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}
