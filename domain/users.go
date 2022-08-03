package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID       uint
	Fullname string
	Username string
	Phone    string
	Password string
	Role     string
}

type UserHandler interface {
	InsertUser() echo.HandlerFunc
	LoginHandler() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}
type UserUsecase interface {
	AddUser(newUser User) (row int, err error)
	LoginUser(userLogin User) (row int, data User, err error)
	UpdateUser(id int, updateProfile User) (row int, err error)
	GetProfile(id int) (User, error)
	DeleteUser(id int) (row int, err error)
}

type UserData interface {
	Insert(newUser User) (row int, err error)
	Login(userLogin User) (row int, data User, err error)
	Update(userID int, updatedData User) (row int, err error)
	GetSpecific(userID int) (User, error)
	Delete(userID int) (row int, err error)
}
