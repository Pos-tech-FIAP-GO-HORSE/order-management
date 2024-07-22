package products

import (
	"context"
	"errors"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/delete_product_by_id"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type DeleteProductByIDUseCase struct {
	ProductRepository repositories.IProductRepository
}

func NewDeleteProductByIDUseCase(productRepository repositories.IProductRepository) delete_product_by_id.IDeleteProductByIDUseCase {
	return &DeleteProductByIDUseCase{
		ProductRepository: productRepository,
	}
}

func (c *DeleteProductByIDUseCase) Execute(ctx context.Context, input delete_product_by_id.Input) error {
	if input.ID == 0 {
		return errors.New("invalid id provided")
	}

	return c.ProductRepository.Delete(ctx, input.ID)
}
