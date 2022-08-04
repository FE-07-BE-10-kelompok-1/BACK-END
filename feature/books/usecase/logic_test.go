package usecase

import (
	"bookstore/config"
	"bookstore/domain"
	"bookstore/domain/mocks"
	"bookstore/infrastructure/aws/s3"
	"errors"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBook(t *testing.T) {
	repo := mocks.BookData{}
	usecase := New(&repo)
	mockBook := domain.Book{ID: 1, Title: "One Piece", Image: "url.com", Price: 20000, Stock: 1, Author: "oda", Sinopsis: "1999", File: "yahahaha"}
	mockInsert := domain.Book{Title: "One Piece", Image: "url.com", Price: 20000, Stock: 1, Author: "oda", Sinopsis: "1999", File: "yahahaha"}

	t.Run("success add new book", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(mockBook, nil).Once()
		res, err := usecase.AddBook(mockInsert)
		assert.Nil(t, err)
		assert.Equal(t, mockBook.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("failed add book", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Book{}, errors.New("cant create book")).Once()
		res, err := usecase.AddBook(mockInsert)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Book{}.ID, res.ID)
		repo.AssertExpectations(t)
	})
}

func TestGetUser(t *testing.T) {
	repo := mocks.BookData{}
	usecase := New(&repo)

	t.Run("success GetUser check admin", func(t *testing.T) {
		repo.On("GetUser", uint(1)).Return(nil).Once()
		err := usecase.GetUser(uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed getUser", func(t *testing.T) {
		repo.On("GetUser", uint(1)).Return(errors.New("youre not admin")).Once()
		err := usecase.GetUser(uint(1))
		assert.EqualError(t, errors.New("youre not admin"), err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := mocks.BookData{}
	usecase := New(&repo)
	mockBook := []domain.Book{{ID: 1, Title: "One Piece", Image: "url.com", Price: 20000, Stock: 1, Author: "oda", Sinopsis: "1999", File: "yahahaha"}}

	t.Run("success getAllBooks", func(t *testing.T) {
		repo.On("GetAll").Return(mockBook, nil).Once()
		res, err := usecase.GetAllBooks()
		assert.Nil(t, err)
		assert.Equal(t, mockBook, res)
		repo.AssertExpectations(t)
	})
}

func TestGetSpecificBook(t *testing.T) {
	repo := mocks.BookData{}
	usecase := New(&repo)
	mockBook := domain.Book{ID: 1, Title: "One Piece", Image: "url.com", Price: 20000, Stock: 1, Author: "oda", Sinopsis: "1999", File: "yahahaha"}

	t.Run("success get specific book", func(t *testing.T) {
		repo.On("GetSpecific", uint(1)).Return(mockBook, nil).Once()
		res, err := usecase.GetSpecificBook(uint(1))
		assert.Nil(t, err)
		assert.Equal(t, res, mockBook)
		repo.AssertExpectations(t)
	})
}

func TestUpdateBook(t *testing.T) {
	repo := mocks.BookData{}
	usecase := New(&repo)
	mockBook := domain.Book{ID: 1, Title: "One Piece", Image: "url.com", Price: 20000, Stock: 1, Author: "oda", Sinopsis: "1999", File: "yahahaha"}

	t.Run("success update book data", func(t *testing.T) {
		repo.On("Update", uint(1), mockBook).Return(mockBook, nil).Once()
		res, err := usecase.UpdateBook(uint(1), mockBook)
		assert.Nil(t, err)
		assert.Greater(t, uint(res.ID), uint(0))
		repo.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	repo := mocks.BookData{}
	usecase := New(&repo)

	t.Run("success delete book", func(t *testing.T) {
		repo.On("Delete", uint(1)).Return(nil)
		err := usecase.DeleteBook(uint(1))
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUploadFile(t *testing.T) {
	imageTrue, _ := os.Open("./files/ERD.jpg")
	fileTrue, _ := os.Open("./files/strukturresimen.pdf")
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
		Size:     int64(imageTrue.Fd()),
	}
	fileTrueCnv := &multipart.FileHeader{
		Filename: fileTrue.Name(),
		Size:     int64(fileTrue.Fd()),
	}

	config := config.GetConfig()
	session := s3.ConnectAws(config)

	repo := mocks.BookData{}
	usecase := New(&repo)

	t.Run("cant upload image to s3", func(t *testing.T) {
		imageUrl, fileUrl, err := usecase.UploadFiles(session, "bucket", imageTrueCnv, fileTrueCnv)
		assert.NotNil(t, err)
		assert.Equal(t, "", imageUrl)
		assert.Equal(t, "", fileUrl)
		repo.AssertExpectations(t)
	})

	t.Run("upload pdf in image", func(t *testing.T) {
		imageUrl, fileUrl, err := usecase.UploadFiles(session, "bucket", fileTrueCnv, fileTrueCnv)
		assert.NotNil(t, err)
		assert.Equal(t, "", imageUrl)
		assert.Equal(t, "", fileUrl)
		repo.AssertExpectations(t)
	})

	t.Run("upload image in pdf", func(t *testing.T) {
		imageUrl, fileUrl, err := usecase.UploadFiles(session, "bucket", imageTrueCnv, imageTrueCnv)
		assert.NotNil(t, err)
		assert.Equal(t, "", imageUrl)
		assert.Equal(t, "", fileUrl)
		repo.AssertExpectations(t)
	})
}
