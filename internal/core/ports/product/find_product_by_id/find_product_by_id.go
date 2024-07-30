package find_product_by_id

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
	ID string `uri:"id" json:"id"`
}

type Output struct {
	Product Product `json:"product"`
}

type IFindProductByID interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
