package find_product_by_category

import (
	"context"
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	IsAvailable bool      `json:"isAvailable"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Input struct {
	Category string `uri:"category"`
}

type Output struct {
	Product Product `json:"product"`
}

type IFindProductByCategory interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
