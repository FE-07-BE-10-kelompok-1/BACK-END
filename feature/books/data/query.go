package data

import (
	"bookstore/domain"
	"errors"

	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.BookData {
	return &bookData{
		db: DB,
	}
}

func (bd *bookData) Insert(newBook domain.Book) (domain.Book, error) {
	converted := ToEntity(newBook)
	err := bd.db.Create(&converted).Error
	if err != nil {
		return domain.Book{}, errors.New("cant create book")
	}

	return converted.ToDomain(), nil
}

func (bd *bookData) GetUser(id uint) error {
	var user domain.User
	err := bd.db.Where("id = ? and role = 'admin'", id).First(&user).Error
	if err != nil {
		return errors.New("you are not admin")
	}

	return nil
}

func (bd *bookData) GetAll() ([]domain.Book, error) {
	var booksData []Book
	err := bd.db.Find(&booksData).Error
	if err != nil {
		return []domain.Book{}, err
	}

	var convert []domain.Book
	for i := 0; i < len(booksData); i++ {
		convert = append(convert, booksData[i].ToDomain())
	}

	return convert, nil
}

func (bd *bookData) GetSpecific(id uint) (domain.Book, error) {
	var book Book
	err := bd.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return domain.Book{}, err
	}

	return book.ToDomain(), nil
}

func (bd *bookData) Update(id uint, updatedData domain.Book) (domain.Book, error) {
	err := bd.db.Model(&Book{}).Where("id = ?", id).Updates(updatedData).Error
	if err != nil {
		return domain.Book{}, err
	}

	var currentData Book
	err = bd.db.Where("id = ?", id).First(&currentData).Error
	if err != nil {
		return domain.Book{}, err
	}

	return currentData.ToDomain(), nil
}

func (bd *bookData) Delete(id uint) error {
	res := bd.db.Where("id = ?", id).Delete(&Book{})
	if res.RowsAffected < 1 {
		return errors.New("no book deleted")
	}

	return nil
}
