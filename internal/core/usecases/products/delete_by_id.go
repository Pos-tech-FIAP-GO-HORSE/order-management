package products

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/delete_product_by_id"
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

func (uc *DeleteProductByIDUseCase) Execute(ctx context.Context, input delete_product_by_id.Input) error {
	_, err := uc.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	return uc.ProductRepository.Delete(ctx, input.ID)
}
