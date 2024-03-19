package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	Name string
	Email string
	Class string
}