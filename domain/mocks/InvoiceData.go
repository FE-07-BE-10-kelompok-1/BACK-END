// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "bookstore/domain"

	mock "github.com/stretchr/testify/mock"
)

// InvoiceData is an autogenerated mock type for the InvoiceData type
type InvoiceData struct {
	mock.Mock
}

// CheckStocks provides a mock function with given fields: _a0
func (_m *InvoiceData) CheckStocks(_a0 []uint) ([]domain.Book, error) {
	ret := _m.Called(_a0)

	var r0 []domain.Book
	if rf, ok := ret.Get(0).(func([]uint) []domain.Book); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCarts provides a mock function with given fields: id
func (_m *InvoiceData) DeleteCarts(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *InvoiceData) GetAll() ([]domain.GetAllInvoices, error) {
	ret := _m.Called()

	var r0 []domain.GetAllInvoices
	if rf, ok := ret.Get(0).(func() []domain.GetAllInvoices); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.GetAllInvoices)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyOrders provides a mock function with given fields: id
func (_m *InvoiceData) GetMyOrders(id uint) ([]domain.GetAllInvoices, error) {
	ret := _m.Called(id)

	var r0 []domain.GetAllInvoices
	if rf, ok := ret.Get(0).(func(uint) []domain.GetAllInvoices); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.GetAllInvoices)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrder provides a mock function with given fields: id, userID
func (_m *InvoiceData) GetOrder(id string, userID uint) error {
	ret := _m.Called(id, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint) error); ok {
		r0 = rf(id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: id
func (_m *InvoiceData) GetUser(id uint) (domain.User, error) {
	ret := _m.Called(id)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(uint) domain.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: data, id
func (_m *InvoiceData) Insert(data domain.Invoice, id []uint) (domain.Invoice, error) {
	ret := _m.Called(data, id)

	var r0 domain.Invoice
	if rf, ok := ret.Get(0).(func(domain.Invoice, []uint) domain.Invoice); ok {
		r0 = rf(data, id)
	} else {
		r0 = ret.Get(0).(domain.Invoice)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Invoice, []uint) error); ok {
		r1 = rf(data, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: data, id
func (_m *InvoiceData) Update(data domain.Invoice, id string) error {
	ret := _m.Called(data, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Invoice, string) error); ok {
		r0 = rf(data, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStock provides a mock function with given fields: _a0
func (_m *InvoiceData) UpdateStock(_a0 []uint) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]uint) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStockAfterCancel provides a mock function with given fields: id
func (_m *InvoiceData) UpdateStockAfterCancel(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewInvoiceData interface {
	mock.TestingT
	Cleanup(func())
}

// NewInvoiceData creates a new instance of InvoiceData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInvoiceData(t mockConstructorTestingTNewInvoiceData) *InvoiceData {
	mock := &InvoiceData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}