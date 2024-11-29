package payment

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

type Payment struct {
	Client payment.Client
}

func NewPaymentClient(client payment.Client) *Payment {
	return &Payment{
		Client: client,
	}
}

type ResponsePayment struct {
	QRCode string `json:"qr_code"`
	ID     int    `json:"id"`
}

func (p *Payment) GeneratePaymentToOrder(ctx context.Context, amount float64, description string, email string) (*ResponsePayment, error) {
	request := payment.Request{
		TransactionAmount: amount,
		Description:       description,
		PaymentMethodID:   "pix",
		Payer: &payment.PayerRequest{
			Email: email,
		},
	}

	response, err := p.Client.Create(ctx, request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	responsePayment := &ResponsePayment{
		QRCode: response.PointOfInteraction.TransactionData.QRCode,
		ID:     response.ID,
	}

	return responsePayment, nil
}
