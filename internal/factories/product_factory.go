package factories

import (
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

func MakeProductFactory(productRepository repositories.IProductRepository) *handlers.ProductHandler {
	createProductUseCase := products.NewCreateProductUseCase(productRepository)
	findAllProductsUseCase := products.NewFindAllProductsUseCase(productRepository)
	findProductByIDUseCase := products.NewFindProductByIDUseCase(productRepository)
	updateProductUseCase := products.NewUpdateProductUseCase(productRepository)
	updateProductAvailabilityUseCase := products.NewUpdateProductAvailabilityUseCase(productRepository)
	deleteProductByIDUseCase := products.NewDeleteProductByIDUseCase(productRepository)

	return handlers.NewProductHandler(
		createProductUseCase,
		findAllProductsUseCase,
		findProductByIDUseCase,
		updateProductUseCase,
		updateProductAvailabilityUseCase,
		deleteProductByIDUseCase,
	)
}
