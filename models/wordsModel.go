package models

import "gorm.io/gorm"

type Words struct {
	gorm.Model
	FirstName string
	LastName  string
}
