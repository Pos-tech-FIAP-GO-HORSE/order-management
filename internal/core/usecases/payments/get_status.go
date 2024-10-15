package payments

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payments_processor"
)

type GetStatusPaymentUseCase struct {
	PaymentService payments_processor.IPaymentProcessor
}

func NewGetStatusPaymentUseCase(paymentService payments_processor.IPaymentProcessor) payment.IGetStatusPaymentUseCase {
	return &GetStatusPaymentUseCase{
		PaymentService: paymentService,
	}
}

func (uc *GetStatusPaymentUseCase) Execute(ctx context.Context, paymentId int) (*payments_processor.ResponseStatusPayment, error) {
	paymentInfos, err := uc.PaymentService.GetPaymentStatus(ctx, paymentId)
	if err != nil {
		return nil, err
	}
	return paymentInfos, nil
}
