package models

import "gorm.io/gorm"

type Students struct {
	gorm.Model
	Name string `json:"name"`
}
