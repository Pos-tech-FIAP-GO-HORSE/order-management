package update_product

import (
	"context"
)

type Input struct {
	ID          string  `json:"id"`
	Name        string  `json:"name,omitempty"`
	Category    string  `json:"category,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Description string  `json:"description,omitempty"`
	ImageUrl    string  `json:"imageUrl,omitempty"`
}

type IUpdateProductUseCase interface {
	Execute(ctx context.Context, input Input) error
}
