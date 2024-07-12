package delete_product_by_id

import (
	"context"
)

type Input struct {
	ID int64 `uri:"id" json:"id"`
}

type IDeleteProductByIDUseCase interface {
	Execute(ctx context.Context, input Input) error
}
