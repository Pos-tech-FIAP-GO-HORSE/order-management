package process_payment

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/orders"
	"time"
)

type Input struct {
	ID                       string             ` json:"id"`
	UserID                   string             ` json:"userId"`
	OrderID                  string             ` json:"orderId"`
	Items                    []*orders.Item     ` json:"items"`
	TotalPrice               float64            ` json:"totalPrice"`
	Status                   orders.OrderStatus ` json:"status"`
	PaymentMethodID          string             ` json:"payment_method_id" validate:"required"`
	EstimatedPreparationTime int64              ` json:"estimatedPreparationTime"`
	CreatedAt                time.Time          ` json:"createdAt"`
	UpdatedAt                time.Time          ` json:"updatedAt"`
	Payer                    PayerInfo          ` json:"payer" validate:"required"`
}

type PayerInfo struct {
	Email string `json:"email" validate:"required,email"` // Email do pagador
}

type IProcessPaymentUseCase interface {
	Execute(ctx context.Context, input Input) error
}
