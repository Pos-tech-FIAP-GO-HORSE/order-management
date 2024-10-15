package payments

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payments_processor"
)

type CreatePaymentUseCase struct {
	PaymentService payments_processor.IPaymentProcessor
}

func NewCreatePaymentUseCase(paymentService payments_processor.IPaymentProcessor) payment.ICreatePaymentUseCase {
	return &CreatePaymentUseCase{
		PaymentService: paymentService,
	}
}

func (uc *CreatePaymentUseCase) Execute(ctx context.Context, input payment.Input) (*payments_processor.ResponseCreatePayment, error) {
	paymentInfos, err := uc.PaymentService.GeneratePaymentToOrder(ctx, input.Amount, input.Description, input.Email)
	if err != nil {
		return nil, err
	}
	return paymentInfos, nil
}
