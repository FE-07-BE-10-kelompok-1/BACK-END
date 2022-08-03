package data

import (
	"bookstore/domain"
	"errors"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.CartData {
	return &cartData{
		db: DB,
	}
}

func (cd *cartData) Insert(data domain.Cart) error {
	err := cd.db.Where("books_id = ? and users_id = ?", data.Books_ID, data.Users_ID).First(&Cart{}).Error
	if err == nil {
		return errors.New("you have added this book to cart before")
	}

	cartData := ToEntity(data)
	err = cd.db.Create(&cartData).Error
	if err != nil {
		return err
	}

	return nil
}

func (cd *cartData) GetAll(id uint) ([]domain.JoinCartWithBooks, error) {
	var cartJoinsData []domain.JoinCartWithBooks
	err := cd.db.Model(&Cart{}).Select("carts.id, carts.books_id, books.title, books.price, books.image").Joins("inner join books on carts.books_id = books.id").Scan(&cartJoinsData).Error
	if err != nil {
		return nil, err
	}
	return cartJoinsData, nil
}

func (cd *cartData) Delete(data domain.Cart) error {
	res := cd.db.Where("id = ? and users_id = ?", data.ID, data.Users_ID).Delete(&Cart{})
	if res.RowsAffected < 1 {
		return errors.New("no data deleted")
	}

	return nil

}
