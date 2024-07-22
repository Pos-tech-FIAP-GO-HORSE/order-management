package products

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/find_product_by_id"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type FindProductByIDUseCase struct {
	ProductRepository repositories.IProductRepository
}

func NewFindProductByIDUseCase(productRepository repositories.IProductRepository) find_product_by_id.IFindProductByID {
	return &FindProductByIDUseCase{
		ProductRepository: productRepository,
	}
}

func (uc *FindProductByIDUseCase) Execute(ctx context.Context, input find_product_by_id.Input) (find_product_by_id.Output, error) {
	product, err := uc.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		return find_product_by_id.Output{}, err
	}

	output := find_product_by_id.Output{
		Product: find_product_by_id.Product{
			ID:          product.ID,
			Name:        product.Name,
			Category:    product.Category,
			Price:       product.Price,
			Description: product.Description,
			ImageUrl:    product.ImageUrl,
			IsAvailable: product.IsAvailable,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		},
	}

	return output, nil
}
