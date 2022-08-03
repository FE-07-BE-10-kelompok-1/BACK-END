package data

import (
	"bookstore/domain"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Books_ID uint
	Users_ID uint
}

func ToEntity(data domain.Cart) Cart {
	return Cart{
		Books_ID: data.Books_ID,
		Users_ID: data.Users_ID,
	}
}

func (c *Cart) ToDomain() domain.Cart {
	return domain.Cart{
		ID:       c.ID,
		Books_ID: c.Books_ID,
		Users_ID: c.Users_ID,
	}
}
