package orders

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/update_order"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
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

	items := make([]*entity.Item, len(input.Items))
	for i, item := range input.Items {
		items[i] = &entity.Item{
			ID:       item.ID,
			Name:     item.Name,
			Price:    item.Price,
			Comments: item.Comments,
		}
	}

	order := &entity.UpdateOrder{
		UserID:                   input.UserID,
		Items:                    items,
		TotalPrice:               input.TotalPrice,
		EstimatedPreparationTime: input.EstimatedPreparationTime,
	}

	return uc.OrderRepository.UpdateByID(ctx, input.ID, order)
}
