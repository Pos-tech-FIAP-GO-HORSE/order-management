package create_order

import (
	"context"
)

type Input struct {
	UserID string `json:"userId"`
	Items  []Item `json:"items"`
}

type Item struct {
	ID       string `json:"id"`
	Comments string `json:"comments"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	OrderID string `json:"orderId"`
}

type ICreateOrderUseCase interface {
	Execute(ctx context.Context, input Input) (string, error)
}
