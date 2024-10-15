package create_payment

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payment"
)

type Input struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Email       string  `json:"email"`
}

type ICreatePaymentUseCase interface {
	Execute(ctx context.Context, input Input) (*payment.ResponsePayment, error)
}
