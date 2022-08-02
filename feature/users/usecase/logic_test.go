package usecase

import (
	"bookstore/domain"
	"bookstore/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.MockUser)
	insertData := domain.User{
		ID:       1,
		Fullname: "Rizuana Nadifatul",
		Username: "rizunadiva",
		Phone:    "081936665965",
		Password: "12345678",
		Role:     "user",
	}
	outputData := domain.User{
		ID:       1,
		Fullname: "Rizuana Nadifatul",
		Username: "rizunadiva",
		Phone:    "081936665965",
		Password: "12345678",
		Role:     "user",
	}
	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(outputData, nil).Once()

		srv := New(repo, validator.New())

		res, err := srv.AddUser(insertData)
		assert.Nil(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})
}
