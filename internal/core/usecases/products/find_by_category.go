package products

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/find_product_by_category"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type FindProductByCategoryUseCase struct {
	ProductRepository repositories.IProductRepository
}

func NewFindProductByCategoryUseCase(productRepository repositories.IProductRepository) find_product_by_category.IFindProductByCategory {
	return &FindProductByCategoryUseCase{
		ProductRepository: productRepository,
	}
}

func (uc *FindProductByCategoryUseCase) Execute(ctx context.Context, input find_product_by_category.Input) (find_product_by_category.Output, error) {
	product, err := uc.ProductRepository.FindByCategory(ctx, input.Category)
	if err != nil {
		return find_product_by_category.Output{}, err
	}

	output := find_product_by_category.Output{
		Product: find_product_by_category.Product{
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
