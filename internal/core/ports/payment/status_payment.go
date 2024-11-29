package payment

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payments_processor"
)

type IGetStatusPaymentUseCase interface {
	Execute(ctx context.Context, paymentId int) (*payments_processor.ResponseStatusPayment, error)
}
