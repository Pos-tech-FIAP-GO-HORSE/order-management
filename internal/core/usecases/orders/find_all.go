package orders

import (
	"context"
	"sort"

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

	orders := make([]find_all_orders.Order, 0, len(foundOrders))

	for _, order := range foundOrders {
		// Only include orders with status 'Ready', 'Preparing', or 'Received'
		if order.Status != "Ready" && order.Status != "Preparing" && order.Status != "Received" {
			continue
		}

		items := make([]find_all_orders.Item, len(order.Items))

		for i := range order.Items {
			items[i] = find_all_orders.Item{
				ID:       order.Items[i].ID,
				Name:     order.Items[i].Name,
				Price:    order.Items[i].Price,
				Comments: order.Items[i].Comments,
			}
		}

		orders = append(orders, find_all_orders.Order{
			ID:                       order.ID,
			UserID:                   order.UserID,
			Items:                    items,
			TotalPrice:               order.TotalPrice,
			EstimatedPreparationTime: order.EstimatedPreparationTime,
			Status:                   string(order.Status),
			CreatedAt:                order.CreatedAt,
			UpdatedAt:                order.UpdatedAt,
		})
	}

	// Sort orders by status priority first, then by date (latest first)
	sort.Slice(orders, func(i, j int) bool {
		statusPriority := map[string]int{"Ready": 0, "Preparing": 1, "Received": 2}

		// First, sort by status priority
		if statusPriority[orders[i].Status] != statusPriority[orders[j].Status] {
			return statusPriority[orders[i].Status] < statusPriority[orders[j].Status]
		}

		// If status is the same, sort by date (latest first)
		return orders[i].CreatedAt.After(orders[j].CreatedAt)
	})

	return find_all_orders.Output{
		CurrentPage: page,
		Orders:      orders,
	}, nil
}
