package find_all_orders

import (
	"context"
	"time"
)

type Order struct {
	ID                       string    `json:"id"`
	UserID                   string    `json:"userId"`
	Items                    []Item    `json:"items"`
	TotalPrice               float64   `json:"totalPrice"`
	EstimatedPreparationTime int64     `json:"estimatedPreparationTime"`
	Status                   string    `json:"status"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
}

type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Comments string  `json:"comments"`
}

type Input struct {
	Page   int64  `form:"page"`
	Limit  int64  `form:"limit"`
	Status string `form:"status"`
}

type Output struct {
	CurrentPage int64   `json:"currentPage"`
	Orders      []Order `json:"orders"`
}

type IFindAllOrdersUseCase interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
