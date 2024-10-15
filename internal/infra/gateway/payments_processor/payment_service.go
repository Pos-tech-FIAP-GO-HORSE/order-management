package payments_processor

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

type ResponseCreatePayment struct {
	QRCode string `json:"qr_code"`
	ID     int    `json:"id"`
}

type ResponseStatusPayment struct {
	ID                int     `json:"id"`
	Status            string  `json:"status"`
	StatusDetail      string  `json:"status_detail"`
	TransactionAmount float64 `json:"transaction_amount"`
	PaymentMethodId   string  `json:"payment_method_id"`
}

func (p *Payment) GeneratePaymentToOrder(ctx context.Context, amount float64, description string, email string) (*ResponseCreatePayment, error) {
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

	responsePayment := &ResponseCreatePayment{
		QRCode: response.PointOfInteraction.TransactionData.QRCode,
		ID:     response.ID,
	}

	return responsePayment, nil
}

func (p *Payment) GetPaymentStatus(ctx context.Context, paymentId int) (*ResponseStatusPayment, error) {
	response, err := p.Client.Get(ctx, paymentId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	responseStatusPayment := &ResponseStatusPayment{
		ID:                response.ID,
		Status:            response.Status,
		StatusDetail:      response.StatusDetail,
		TransactionAmount: response.TransactionAmount,
		PaymentMethodId:   response.PaymentMethodID,
	}

	return responseStatusPayment, nil
}
