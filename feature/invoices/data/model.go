package data

import (
	"bookstore/feature/orders/data"
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ID             string `gorm:"type:VARCHAR(255);primarykey"`
	Users_ID       uint
	Total          int
	Payment_Link   string
	Payment_Method string
	Paid_At        time.Time
	Orders         []data.Order `gorm:"foreignKey:Invoice_ID"`
}
