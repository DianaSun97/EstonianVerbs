package models

import "gorm.io/gorm"

type Words struct {
	gorm.Model
	EstVerb string
	EngVerb string
}
