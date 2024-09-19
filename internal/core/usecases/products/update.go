package products

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
	valueobjects "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/valueObjects"
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
	category, err := valueobjects.ParseToProductCategoryType(input.Category)
	if err != nil {
		return err
	}

	_, err = uc.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	product := &entity.UpdateProduct{
		Name:            input.Name,
		Category:        category,
		Price:           input.Price,
		Description:     input.Description,
		ImageUrl:        input.ImageUrl,
		PreparationTime: input.PreparationTime,
	}

	return uc.ProductRepository.UpdateByID(ctx, input.ID, product)
}
