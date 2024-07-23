package products

import (
	"context"
	domain_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/models"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/create_product"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type CreateProductUseCase struct {
	ProductRepository repositories.IProductRepository
}

func NewCreateProductUseCase(productRepository repositories.IProductRepository) create_product.ICreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (c *CreateProductUseCase) Execute(ctx context.Context, input create_product.Input) error {

	product := &domain_products.Product{
		Name:        input.Name,
		Category:    input.Category,
		Price:       input.Price,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		IsAvailable: input.IsAvailable,
	}

	if err := c.ProductRepository.Create(ctx, product); err != nil {
		return err
	}

	return nil
}
