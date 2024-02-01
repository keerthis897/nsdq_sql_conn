package models

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	ID int `json:"id"`

	Name string `json:"name"`
}
