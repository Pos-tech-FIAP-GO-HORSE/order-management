package payments_processor

import (
	"context"
)

type IPaymentProcessor interface {
	GeneratePaymentToOrder(ctx context.Context, amount float64, description string, email string) (*ResponseCreatePayment, error)
	GetPaymentStatus(ctx context.Context, paymentId int) (*ResponseStatusPayment, error)
}
