package delivery

type CheckoutReq struct {
	Books_ID []uint `json:"books_id" form:"books_id" validate:"required"`
}

type MidtransCallbackRequest struct {
	Order_ID           string `json:"order_id" form:"order_id"`
	Transaction_Status string `json:"transaction_status" form:"transaction_status"`
	Settlement_Time    string `json:"settlement_time" form:"settlement_time"`
	Payment_Type       string `json:"payment_type" form:"payment_type"`
	Gross_Amount       string `json:"gross_amount" form:"gross_amount"`
}
