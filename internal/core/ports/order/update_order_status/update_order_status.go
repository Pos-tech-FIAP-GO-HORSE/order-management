package update_order_status

import (
	"context"
)

type Input struct {
	ID     string `uri:"id" swaggerignore:"true"`
	Status string `json:"status" enums:"Received,Awaiting Payment,Confirmed,Preparing,Ready,Finished,Canceled"`
}

type IUpdateOrderStatusUseCase interface {
	Execute(ctx context.Context, input Input) error
}
