package data

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Books_ID uint
	Users_ID uint
}
