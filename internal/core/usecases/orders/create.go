package orders

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
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

func (uc *CreateOrderUseCase) Execute(ctx context.Context, input create_order.Input) (create_order.Output, error) {
	items := make([]*entity.Item, 0)

	if input.UserID != "" {
		_, err := uc.UserRepository.FindByID(ctx, input.UserID)
		if err != nil {
			return create_order.Output{}, err
		}
	}

	for _, item := range input.Items {
		product, err := uc.ProductRepository.FindByID(ctx, item.ID)
		if err != nil {
			return create_order.Output{}, err
		}

		newItem, err := entity.NewItem(item.ID, product.Name, item.Comments, product.Price, product.PreparationTime)
		if err != nil {
			return create_order.Output{}, err
		}

		items = append(items, newItem)
	}

	newOrder, err := entity.NewOrder(input.UserID, items)
	if err != nil {
		return create_order.Output{}, err
	}

	newOrder.CalculateTotalPrice()
	newOrder.CalculateEstimatedPreparationTime()

	orderID, err := uc.OrderRepository.Create(ctx, newOrder)
	if err != nil {
		return create_order.Output{}, err
	}
	return create_order.Output{ID: orderID.ID}, nil
}
