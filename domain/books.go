package domain

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws/session"
)

type Book struct {
	ID       uint   `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Image    string `json:"image" form:"image"`
	Price    int    `json:"price" form:"price"`
	Stock    int    `json:"stock" form:"stock"`
	Author   string `json:"author" form:"author"`
	Sinopsis string `json:"sinopsis" form:"sinopsis"`
	File     string `json:"file" form:"file"`
}

type BookUsecase interface {
	AddBook(newBook Book) (Book, error)
	UploadFiles(session *session.Session, bucket string, image *multipart.FileHeader, file *multipart.FileHeader) (string, string, error)
	GetUser(id uint) error
	GetAllBooks() ([]Book, error)
	GetSpecificBook(id uint) (Book, error)
	UpdateBook(id uint, updatedData Book) (Book, error)
	DeleteBook(id uint) error
}

type BookData interface {
	Insert(newBook Book) (Book, error)
	GetUser(id uint) error
	GetAll() ([]Book, error)
	GetSpecific(id uint) (Book, error)
	Update(id uint, updatedData Book) (Book, error)
	Delete(id uint) error
}
