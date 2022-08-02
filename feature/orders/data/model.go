package data

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Invoice_ID string `gorm:"type:VARCHAR(255)"`
	Books_ID   uint
}
