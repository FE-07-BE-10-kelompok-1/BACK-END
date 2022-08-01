package delivery

import "bookstore/domain"

type UserResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func FromModel(data domain.User) UserResponse {
	var res UserResponse
	res.ID = data.ID
	res.Fullname = data.Fullname
	res.Username = data.Username
	res.Phone = data.Phone
	res.Password = data.Password
	res.Role = data.Role
	return res
}
