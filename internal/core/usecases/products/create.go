package products

import (
	"context"

	domain_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/create_product"
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

func (uc *CreateProductUseCase) Execute(ctx context.Context, input create_product.Input) error {
	product, err := domain_products.NewProduct(input.Name, input.Category, input.Description, input.ImageUrl, input.Price)
	if err != nil {
		return err
	}

	return uc.ProductRepository.Create(ctx, product)
}
