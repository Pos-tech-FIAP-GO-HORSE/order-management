package payment

import (
	"context"
	"errors"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/payment/process_payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

// MercadoPagoClient define a interface para se comunicar com a API do Mercado Pago
type MercadoPagoClient interface {
	CreatePayment(input *process_payment.Input) (string, error) // retorna ID do pagamento ou erro
}

// ProcessPaymentUseCase é a implementação do serviço de pagamento
type ProcessPaymentUseCase struct {
	PaymentRepository repositories.IPaymentRepository
	mercadoPagoClient MercadoPagoClient
}

func NewProcessPaymentUseCase(paymentRepository repositories.IPaymentRepository, mercadoPagoClient MercadoPagoClient) process_payment.IProcessPaymentUseCase {
	return &ProcessPaymentUseCase{
		PaymentRepository: paymentRepository,
		mercadoPagoClient: mercadoPagoClient,
	}
}

func (s *ProcessPaymentUseCase) Execute(ctx context.Context, input process_payment.Input) error {
	// Realizar a comunicação com a API do Mercado Pago
	paymentID, err := s.mercadoPagoClient.CreatePayment(&input)
	if err != nil {
		return errors.New("falha ao processar o pagamento: " + err.Error())
	}

	// Criar a entidade Payment
	paymentProcessed := &payment.Payment{
		OrderID:   input.ID,
		Amount:    input.TotalPrice,
		PaymentID: paymentID,
		Status:    "Pending", // ou qualquer outro status inicial
	}

	// Salvar o pagamento no banco de dados
	err = s.PaymentRepository.Save(ctx, paymentProcessed)
	if err != nil {
		return errors.New("falha ao salvar o pagamento: " + err.Error())
	}

	return nil
}
