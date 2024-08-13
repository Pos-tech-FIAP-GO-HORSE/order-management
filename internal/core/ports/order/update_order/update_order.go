package update_order

import (
	"context"
)

type Input struct {
	ID                       string  `uri:"id" swaggerignore:"true"`
	UserID                   string  `json:"userId"`
	Items                    []Item  `json:"items"`
	TotalPrice               float64 `json:"totalPrice"`
	EstimatedPreparationTime int64   `json:"estimatedPreparationTime"`
	Status                   string  `json:"status" enums:"Received,Awaiting Payment,Confirmed,Preparing,Ready,Finished,Canceled"`
}

type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Comments string  `json:"comments"`
}

type IUpdateOrderUseCase interface {
	Execute(ctx context.Context, input Input) error
}
