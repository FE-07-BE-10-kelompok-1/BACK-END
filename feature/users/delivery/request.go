package delivery

import "bookstore/domain"

type InsertFormat struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required,min=10"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Fullname: i.Fullname,
		Username: i.Username,
		Phone:    i.Phone,
		Password: i.Password,
	}
}

type LoginFormat struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (lf *LoginFormat) LoginToModel() domain.User {
	return domain.User{
		Username: lf.Username,
		Password: lf.Password,
	}
}
