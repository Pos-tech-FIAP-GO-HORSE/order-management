package products

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
	values "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/valueObjects"
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

func (uc *CreateProductUseCase) Execute(ctx context.Context, input create_product.Input) error {
	category, err := values.ParseToProductCategoryType(input.Category)
	if err != nil {
		return err
	}

	product, err := entity.NewProduct(input.Name, input.Description, input.ImageUrl, category, input.Price, input.PreparationTime)
	if err != nil {
		return err
	}

	return uc.ProductRepository.Create(ctx, product)
}
