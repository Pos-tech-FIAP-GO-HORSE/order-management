package orders

import (
	"context"

	valueobjects "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/valueObjects"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/update_order_status"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UpdateOrderStatusUseCase struct {
	OrderRepository repositories.IOrderRepository
}

func NewUpdateOrderStatusUseCase(orderRepository repositories.IOrderRepository) update_order_status.IUpdateOrderStatusUseCase {
	return &UpdateOrderStatusUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *UpdateOrderStatusUseCase) Execute(ctx context.Context, input update_order_status.Input) error {
	status, err := valueobjects.ParseToOrderStatusType(input.Status)
	if err != nil {
		return err
	}

	_, err = uc.OrderRepository.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	return uc.OrderRepository.UpdateStatusByID(ctx, input.ID, status)
}
