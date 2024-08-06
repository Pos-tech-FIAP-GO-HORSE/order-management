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

type ICreateOrderUseCase interface {
	Execute(ctx context.Context, input Input) error
}
