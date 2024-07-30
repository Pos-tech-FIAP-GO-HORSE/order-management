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
	foundProducts, err := uc.ProductRepository.FindByCategory(ctx, input.Category)
	if err != nil {
		return find_product_by_category.Output{}, err
	}

	products := make([]find_product_by_category.Product, 0)

	for _, p := range foundProducts {
		product := find_product_by_category.Product{
			ID:          p.ID,
			Name:        p.Name,
			Category:    p.Category,
			Price:       p.Price,
			Description: p.Description,
			ImageUrl:    p.ImageUrl,
			IsAvailable: p.IsAvailable,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		}

		products = append(products, product)
	}

	return find_product_by_category.Output{
		Products: products,
	}, nil
}
