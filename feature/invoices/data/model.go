package data

import (
	"bookstore/domain"
	"bookstore/feature/orders/data"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ID             string `gorm:"type:VARCHAR(255);primarykey"`
	Users_ID       uint
	Total          int
	Status         string `gorm:"default:waiting"`
	Payment_Link   string
	Payment_Method string       `gorm:"default:NULL"`
	Paid_At        string       `gorm:"default:NULL"`
	Orders         []data.Order `gorm:"foreignKey:Invoice_ID"`
}

func ToEntity(data domain.Invoice) Invoice {
	return Invoice{
		ID:             data.ID,
		Users_ID:       data.Users_ID,
		Total:          data.Total,
		Status:         data.Status,
		Payment_Link:   data.Payment_Link,
		Payment_Method: data.Payment_Method,
		Paid_At:        data.Paid_At,
	}
}

func (i *Invoice) ToDomain() domain.Invoice {
	return domain.Invoice{
		ID:             i.ID,
		Users_ID:       i.Users_ID,
		Total:          i.Total,
		Status:         i.Status,
		Payment_Link:   i.Payment_Link,
		Payment_Method: i.Payment_Method,
		Paid_At:        i.Paid_At,
	}
}
