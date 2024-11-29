package payment

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payments_processor"
)

type Input struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Email       string  `json:"email"`
}

type ICreatePaymentUseCase interface {
	Execute(ctx context.Context, input Input) (*payments_processor.ResponseCreatePayment, error)
}
