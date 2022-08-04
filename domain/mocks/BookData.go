// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "bookstore/domain"

	mock "github.com/stretchr/testify/mock"
)

// BookData is an autogenerated mock type for the BookData type
type BookData struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *BookData) Delete(id uint) error {
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
func (_m *BookData) GetAll() ([]domain.Book, error) {
	ret := _m.Called()

	var r0 []domain.Book
	if rf, ok := ret.Get(0).(func() []domain.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Book)
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

// GetSpecific provides a mock function with given fields: id
func (_m *BookData) GetSpecific(id uint) (domain.Book, error) {
	ret := _m.Called(id)

	var r0 domain.Book
	if rf, ok := ret.Get(0).(func(uint) domain.Book); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: id
func (_m *BookData) GetUser(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: newBook
func (_m *BookData) Insert(newBook domain.Book) (domain.Book, error) {
	ret := _m.Called(newBook)

	var r0 domain.Book
	if rf, ok := ret.Get(0).(func(domain.Book) domain.Book); ok {
		r0 = rf(newBook)
	} else {
		r0 = ret.Get(0).(domain.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Book) error); ok {
		r1 = rf(newBook)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updatedData
func (_m *BookData) Update(id uint, updatedData domain.Book) (domain.Book, error) {
	ret := _m.Called(id, updatedData)

	var r0 domain.Book
	if rf, ok := ret.Get(0).(func(uint, domain.Book) domain.Book); ok {
		r0 = rf(id, updatedData)
	} else {
		r0 = ret.Get(0).(domain.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, domain.Book) error); ok {
		r1 = rf(id, updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookData interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookData creates a new instance of BookData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookData(t mockConstructorTestingTNewBookData) *BookData {
	mock := &BookData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
