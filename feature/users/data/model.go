package data

import (
	"bookstore/domain"
	cartData "bookstore/feature/carts/data"
	"bookstore/feature/invoices/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string
	Username string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Password string
	Role     string          `gorm:"default:user"`
	Carts    []cartData.Cart `gorm:"foreignKey:Users_ID"`
	Invoices []data.Invoice  `gorm:"foreignKey:Users_ID"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:       u.ID,
		Fullname: u.Fullname,
		Username: u.Username,
		Phone:    u.Phone,
		Password: u.Password,
		Role:     u.Role,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Username = data.Username
	res.Password = data.Password
	res.Fullname = data.Fullname
	res.Phone = data.Phone
	return res
}
