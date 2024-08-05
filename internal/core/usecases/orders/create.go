package orders

import (
	"context"

	domain_orders "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/create_order"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderUseCase struct {
	OrderRepository   repositories.IOrderRepository
	ProductRepository repositories.IProductRepository
}

func NewCreateProductUseCase(orderRepository repositories.IOrderRepository, productRepository repositories.IProductRepository) create_order.ICreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository:   orderRepository,
		ProductRepository: productRepository,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input create_order.Input) error {
	items := make([]*domain_orders.Item, 0)

	for _, item := range input.Items {
		_, err := primitive.ObjectIDFromHex(item.ID)
		if err != nil {
			return err
		}

		product, err := uc.ProductRepository.FindByID(ctx, item.ID)
		if err != nil {
			return err
		}

		item, err := domain_orders.NewItem(item.ID, product.Name, item.Comments, product.Price)
		if err != nil {
			return err
		}

		items = append(items, item)
	}

	order, err := domain_orders.NewOrder(input.UserID, items)
	if err != nil {
		return err
	}

	return uc.OrderRepository.Create(ctx, order)
}
