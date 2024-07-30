package update_product_availability

import (
	"context"
)

type Input struct {
	ID string `uri:"id" json:"id"`
}

type IUpdateProductAvailabilityUseCase interface {
	Execute(ctx context.Context, input Input) error
}
