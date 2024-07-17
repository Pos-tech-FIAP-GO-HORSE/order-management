package products

import (
	"context"

	domain_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/update_product"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UpdateProductUseCase struct {
	ProductRepository repositories.IProductRepository
}

func NewUpdateProductUseCase(productRepository repositories.IProductRepository) update_product.IUpdateProductUseCase {
	return &UpdateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (c *UpdateProductUseCase) Execute(ctx context.Context, input update_product.Input) error {
	product := &domain_products.Product{
		Name:        input.Name,
		Category:    input.Category,
		Price:       input.Price,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		IsAvailable: input.IsAvailable,
	}

	if err := c.ProductRepository.Update(ctx, input.ID, product); err != nil {
		return err
	}

	return nil
}
