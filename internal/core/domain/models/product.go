package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `db:"name"`
	Category    string  `db:"category"`
	Price       float64 `db:"price"`
	Description string  `db:"description"`
	ImageUrl    string  `db:"image_url"`
	IsAvailable bool    `db:"is_available"`
}
