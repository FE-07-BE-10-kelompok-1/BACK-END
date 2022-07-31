package data

import (
	"bookstore/feature/orders/data"
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Users_ID uint
	Total    int
	Paid_At  time.Time
	Orders   []data.Order `gorm:"foreignKey:Invoice_ID"`
}
