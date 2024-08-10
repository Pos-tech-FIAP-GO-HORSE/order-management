package products

import (
	"context"

	domain_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/update_product"
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

func (uc *UpdateProductUseCase) Execute(ctx context.Context, input update_product.Input) error {
	_, err := uc.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	product := &domain_products.UpdateProduct{
		Name:            input.Name,
		Category:        input.Category,
		Price:           input.Price,
		Description:     input.Description,
		ImageUrl:        input.ImageUrl,
		PreparationTime: input.PreparationTime,
	}

	return uc.ProductRepository.UpdateByID(ctx, input.ID, product)
}
