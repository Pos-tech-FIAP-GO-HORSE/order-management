package orders

import (
	"context"

	domain_orders "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/create_order"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type CreateOrderUseCase struct {
	OrderRepository   repositories.IOrderRepository
	ProductRepository repositories.IProductRepository
	UserRepository    repositories.IUserRepository
}

func NewCreateProductUseCase(orderRepository repositories.IOrderRepository, productRepository repositories.IProductRepository, userRepository repositories.IUserRepository) create_order.ICreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository:   orderRepository,
		ProductRepository: productRepository,
		UserRepository:    userRepository,
	}
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input create_order.Input) error {
	items := make([]*domain_orders.Item, 0)

	if input.UserID != "" {
		_, err := uc.UserRepository.FindByID(ctx, input.UserID)
		if err != nil {
			return err
		}
	}

	for _, item := range input.Items {
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
