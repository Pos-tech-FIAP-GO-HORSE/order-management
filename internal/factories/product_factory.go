package factories

import (
	product_usecase "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/usecases/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/handlers"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

func MakeProductFactory(productRepository repositories.IProductRepository) *handlers.ProductHandler {
	createProductUseCase := product_usecase.NewCreateProductUseCase(productRepository)
	findAllProductsUseCase := product_usecase.NewFindAllProductsUseCase(productRepository)
	findProductByIDUseCase := product_usecase.NewFindProductByIDUseCase(productRepository)
	updateProductUseCase := product_usecase.NewUpdateProductUseCase(productRepository)
	deleteProductByIDUseCase := product_usecase.NewDeleteProductByIDUseCase(productRepository)

	return handlers.NewProductHandler(
		createProductUseCase,
		findAllProductsUseCase,
		findProductByIDUseCase,
		updateProductUseCase,
		deleteProductByIDUseCase,
	)
}
