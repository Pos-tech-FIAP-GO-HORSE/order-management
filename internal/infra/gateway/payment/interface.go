package payment

import (
	"context"
)

type IPaymentProcessor interface {
	GeneratePaymentToOrder(ctx context.Context, amount float64, description string, email string) (*ResponsePayment, error)
}
