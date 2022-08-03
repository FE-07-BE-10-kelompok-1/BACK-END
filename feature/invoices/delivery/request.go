package delivery

type CheckoutReq struct {
	Books_ID []uint `json:"books_id" form:"books_id" validate:"required"`
}
