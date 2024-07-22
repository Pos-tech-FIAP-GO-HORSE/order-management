package products

import (
	"context"
	find_all_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/find_all_products"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type FindAllProductsUseCase struct {
	ProductRepository repositories.IProductRepository
}

func NewFindAllProductsUseCase(productRepository repositories.IProductRepository) find_all_products.IFindAllProducts {
	return &FindAllProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (f *FindAllProductsUseCase) Execute(ctx context.Context, input find_all_products.Input) (find_all_products.Output, error) {
	page, limit := normalizePage(input.Page), normalizeLimit(input.Limit)
	offset := calculateOffset(page, limit)

	foundProducts, err := f.ProductRepository.Find(ctx, offset, limit)
	if err != nil {
		return find_all_products.Output{}, err
	}

	products := make([]find_all_products.Product, 0)

	for _, p := range foundProducts {
		product := find_all_products.Product{
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

	return find_all_products.Output{
		CurrentPage: page,
		Products:    products,
	}, nil
}
