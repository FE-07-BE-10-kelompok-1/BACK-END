package data

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Invoice_ID uint
	Books_ID   uint
}
