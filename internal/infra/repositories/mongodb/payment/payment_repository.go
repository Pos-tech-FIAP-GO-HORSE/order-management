package payment

import (
	"context"
	"errors"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/payment"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// PaymentRepository é a implementação de PaymentRepository usando MongoDB
type PaymentRepository struct {
	collection *mongo.Collection
}

// NewProcessPaymentRepository cria uma nova instância de PaymentRepository
func NewProcessPaymentRepository(collection *mongo.Collection) repositories.IPaymentRepository {
	return &PaymentRepository{
		collection: collection,
	}
}

// Save salva um novo pagamento no MongoDB
func (r *PaymentRepository) Save(ctx context.Context, payment *payment.Payment) error {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, payment)
	return err
}

// GetByID busca um pagamento pelo ID no MongoDB
func (r *PaymentRepository) GetByID(ctx context.Context, id string) (*payment.Payment, error) {
	var paymentProcessed payment.Payment
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&paymentProcessed)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("paymentProcessed not found")
	}
	return &paymentProcessed, err
}

// Update atualiza um pagamento no MongoDB
func (r *PaymentRepository) Update(ctx context.Context, payment *payment.Payment) error {
	payment.UpdatedAt = time.Now()

	filter := bson.M{"_id": payment.ID}
	update := bson.M{"$set": payment}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
