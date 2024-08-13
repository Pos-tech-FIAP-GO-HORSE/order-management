package orders

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/update_order"

	domain_orders "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UpdateOrderUseCase struct {
	OrderRepository repositories.IOrderRepository
}

func NewUpdateOrderUseCase(orderRepository repositories.IOrderRepository) update_order.IUpdateOrderUseCase {
	return &UpdateOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *UpdateOrderUseCase) Execute(ctx context.Context, input update_order.Input) error {
	_, err := uc.OrderRepository.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	items := make([]*domain_orders.Item, len(input.Items))
	for i, item := range input.Items {
		items[i] = &domain_orders.Item{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price,
			Comments: item.Comments,
		}
	}

	order := &domain_orders.UpdateOrder{
		UserID:                   input.UserID,
		Items:                    items,
		TotalPrice:               input.TotalPrice,
		EstimatedPreparationTime: input.EstimatedPreparationTime,
		Status:                   domain_orders.OrderStatus(input.Status),
	}

	return uc.OrderRepository.UpdateByID(ctx, input.ID, order)
}
