package models

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model
	Name    string `json:"name"`
	HowMany int    `json:"how_many"`
	Books   []Book `json:"books"`
}
