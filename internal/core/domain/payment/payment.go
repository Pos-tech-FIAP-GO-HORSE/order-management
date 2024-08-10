package payment

import "time"

// Payment representa o pagamento de um pedido
type Payment struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	PaymentID string    `json:"payment_id" bson:"payment_id"`
	OrderID   string    `json:"order_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
