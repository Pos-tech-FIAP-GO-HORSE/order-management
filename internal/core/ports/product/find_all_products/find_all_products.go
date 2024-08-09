package find_all_products

import (
	"context"
	"time"
)

type Product struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Category        string    `json:"category"`
	Price           float64   `json:"price"`
	Description     string    `json:"description"`
	ImageUrl        string    `json:"imageUrl"`
	IsAvailable     bool      `json:"isAvailable"`
	PreparationTime int64     `json:"preparationTime"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type Input struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

type Output struct {
	CurrentPage int64     `json:"currentPage"`
	Products    []Product `json:"products"`
}

type IFindAllProducts interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
