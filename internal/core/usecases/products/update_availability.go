package products

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/product/update_product_availability"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UpdateProductAvailabilityUseCase struct {
	ProductRepository repositories.IProductRepository
}

func NewUpdateProductAvailabilityUseCase(productRepository repositories.IProductRepository) update_product_availability.IUpdateProductAvailabilityUseCase {
	return &UpdateProductAvailabilityUseCase{
		ProductRepository: productRepository,
	}
}

func (uc *UpdateProductAvailabilityUseCase) Execute(ctx context.Context, input update_product_availability.Input) error {
	product, err := uc.ProductRepository.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	availability := !product.IsAvailable

	return uc.ProductRepository.UpdateAvailability(ctx, input.ID, availability)
}
