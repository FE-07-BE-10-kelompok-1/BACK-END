package domain

import "time"

type Invoice struct {
	ID       uint
	Users_ID uint
	Total    int
	Paid_At  time.Time
}

type InvoiceUsecase interface{}

type InvoiceData interface{}
