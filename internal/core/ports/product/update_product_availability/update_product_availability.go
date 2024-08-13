package update_product_availability

import (
	"context"
)

type Input struct {
	ID string `uri:"id" swaggerignore:"true"`
}

type IUpdateProductAvailabilityUseCase interface {
	Execute(ctx context.Context, input Input) error
}
