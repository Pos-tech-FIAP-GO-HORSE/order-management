package create_product

import (
	"context"
)

type Input struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"imageUrl"`
	IsAvailable bool    `json:"isAvailable"`
}

type ICreateProductUseCase interface {
	Execute(ctx context.Context, input Input) error
}
