package delivery

import "bookstore/domain"

type Carts struct {
	ID       uint   `json:"id" form:"id"`
	Books_ID uint   `json:"books_id" form:"books_id"`
	Title    string `json:"title" form:"title"`
	Price    string `json:"price" form:"price"`
	Image    string `json:"image" form:"image"`
}

func ToCartsResponse(data []domain.JoinCartWithBooks) []Carts {
	var cartData []Carts
	for i := 0; i < len(data); i++ {
		cartData = append(cartData, Carts{
			ID:       data[i].ID,
			Books_ID: data[i].Books_ID,
			Title:    data[i].Title,
			Price:    data[i].Price,
			Image:    data[i].Image,
		})
	}
	return cartData
}
