package domain

type User struct {
	ID       uint
	Fullname string
	Username string
	Phone    string
	Password string
	Role     string
}

type UserUsecase interface{}

type UserData interface{}
