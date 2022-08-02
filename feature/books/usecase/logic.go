package usecase

import (
	"bookstore/domain"
	"bookstore/infrastructure/aws/s3"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/uuid"
)

type bookUsecase struct {
	bookData domain.BookData
}

func New(bd domain.BookData) domain.BookUsecase {
	return &bookUsecase{
		bookData: bd,
	}
}

func (bs *bookUsecase) AddBook(newBook domain.Book) (domain.Book, error) {
	book, err := bs.bookData.Insert(newBook)
	return book, err
}

func (bs *bookUsecase) GetUser(id uint) error {
	err := bs.bookData.GetUser(id)
	return err
}

func (bs *bookUsecase) GetAllBooks() ([]domain.Book, error) {
	data, err := bs.bookData.GetAll()
	return data, err
}

func (bs *bookUsecase) GetSpecificBook(id uint) (domain.Book, error) {
	data, err := bs.bookData.GetSpecific(id)
	return data, err
}

func (bs *bookUsecase) UpdateBook(id uint, updatedData domain.Book) (domain.Book, error) {
	data, err := bs.bookData.Update(id, updatedData)
	return data, err
}

func (bs *bookUsecase) DeleteBook(id uint) error {
	err := bs.bookData.Delete(id)
	return err
}

func (bs *bookUsecase) UploadFiles(session *session.Session, bucket string, image *multipart.FileHeader, file *multipart.FileHeader) (string, string, error) {
	imageExt := strings.Split(image.Filename, ".")
	ext := imageExt[len(imageExt)-1]
	if ext != "png" && ext != "jpeg" && ext != "jpg" {
		return "", "", errors.New("image not supported, supported: png/jpeg/jpg")
	}
	fileExt := strings.Split(file.Filename, ".")
	ext = fileExt[len(fileExt)-1]
	if ext != "pdf" {
		return "", "", errors.New("file not supported, supported: pdf")
	}

	destination := fmt.Sprint("images/", uuid.NewString(), "_", image.Filename)
	imageUrl, err := s3.DoUpload(session, *image, bucket, destination)
	if err != nil {
		return "", "", errors.New("cant upload image to s3")
	}
	destination = fmt.Sprint("files/", uuid.NewString(), "_", file.Filename)
	fileUrl, err := s3.DoUpload(session, *image, bucket, destination)
	if err != nil {
		return "", "", errors.New("cant upload file to s3")
	}

	return imageUrl, fileUrl, nil
}
