package orders

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/order/find_all_orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/utils"
)

type FindAllOrdersUseCase struct {
	OrderRepository repositories.IOrderRepository
}

func NewFindAllOrdersUseCase(orderRepository repositories.IOrderRepository) find_all_orders.IFindAllOrdersUseCase {
	return &FindAllOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *FindAllOrdersUseCase) Execute(ctx context.Context, input find_all_orders.Input) (find_all_orders.Output, error) {
	page, limit := utils.NormalizePage(input.Page), utils.NormalizeLimit(input.Limit)
	offset := utils.CalculateOffset(page, limit)

	filter := utils.OrderFilters{
		Status: input.Status,
	}

	foundOrders, err := uc.OrderRepository.Find(ctx, filter, offset, limit)
	if err != nil {
		return find_all_orders.Output{}, err
	}

	orders := make([]find_all_orders.Order, len(foundOrders))

	for i, order := range foundOrders {
		items := make([]find_all_orders.Item, len(order.Items))

		for i := range order.Items {
			items[i] = find_all_orders.Item{
				ID:       order.Items[i].ID,
				Name:     order.Items[i].Name,
				Price:    order.Items[i].Price,
				Comments: order.Items[i].Comments,
			}
		}

		orders[i] = find_all_orders.Order{
			ID:         order.ID,
			UserID:     order.UserID,
			Items:      items,
			TotalPrice: order.TotalPrice,
			Status:     string(order.Status),
			CreatedAt:  order.CreatedAt,
			UpdatedAt:  order.UpdatedAt,
		}
	}

	return find_all_orders.Output{
		CurrentPage: page,
		Orders:      orders,
	}, nil
}
