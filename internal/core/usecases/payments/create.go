package payments

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/create_payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/gateway/payment"
)

type CreatePaymentUseCase struct {
	PaymentService payment.IPaymentProcessor
}

func NewCreatePaymentUseCase(paymentService payment.IPaymentProcessor) create_payment.ICreatePaymentUseCase {
	return &CreatePaymentUseCase{
		PaymentService: paymentService,
	}
}

func (uc *CreatePaymentUseCase) Execute(ctx context.Context, input create_payment.Input) (*payment.ResponsePayment, error) {
	paymentInfos, err := uc.PaymentService.GeneratePaymentToOrder(ctx, input.Amount, input.Description, input.Email)
	if err != nil {
		return nil, err
	}
	return paymentInfos, nil
}
