package payment

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/payment/process_payment"
	"net/http"
)

// MercadoPagoClientImpl é a implementação da interface MercadoPagoClient
type MercadoPagoClientImpl struct {
	accessToken string
	baseURL     string
}
type paymentResponse struct {
	ID string `json:"id"`
}

// NewMercadoPagoClient cria uma nova instância de MercadoPagoClientImpl
func NewMercadoPagoClient(accessToken, baseURL string) *MercadoPagoClientImpl {
	return &MercadoPagoClientImpl{
		accessToken: accessToken,
		baseURL:     baseURL,
	}
}

// CreatePayment cria um pagamento via API do Mercado Pago
func (c *MercadoPagoClientImpl) CreatePayment(input *process_payment.Input) (string, error) {
	// Construir a payload para a API do Mercado Pago
	paymentRequest := map[string]interface{}{
		"transaction_amount": input.TotalPrice,
		"installments":       1,
		"description":        "Payment for Order #" + input.ID,
		"payment_method_id":  "credit_card", // ou outro método
		"payer": map[string]interface{}{
			"email": "customer@example.com", // Email do cliente
		},
	}

	payload, err := json.Marshal(paymentRequest)
	if err != nil {
		return "", err
	}

	// Enviar a requisição HTTP para a API do Mercado Pago
	req, err := http.NewRequest("POST", c.baseURL+"/v1/payments", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", errors.New("falha ao criar pagamento no Mercado Pago")
	}

	paymentResponse := paymentResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return "", err
	}

	return paymentResponse.ID, nil
}
