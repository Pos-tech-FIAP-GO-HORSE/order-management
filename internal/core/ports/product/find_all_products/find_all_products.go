package findallproducts

import (
	"context"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	IsAvailable bool    `json:"isAvailable"`
}

type Input struct {
	Page  int64 `form:"page" json:"page"`
	Limit int64 `form:"limit" json:"limit"`
}

type Output struct {
	CurrentPage int64     `json:"currentPage"`
	Products    []Product `json:"products"`
}

type IFindAllProducts interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
