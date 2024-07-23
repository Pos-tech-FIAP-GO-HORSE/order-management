package findproductbyid

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
	ID int64 `uri:"id" json:"id"`
}

type Output struct {
	gorm.Model
	Product Product `json:"product"`
}

type IFindProductByID interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
