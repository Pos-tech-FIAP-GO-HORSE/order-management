package orders

import (
	"context"

	domain_orders "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) repositories.IOrderRepository {
	return &OrderRepository{
		collection: collection,
	}
}

func (o *OrderRepository) Create(ctx context.Context, order *domain_orders.Order) error {
	_, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) Find(ctx context.Context, offset int64, limit int64) ([]*domain_orders.Order, error) {
	panic("unimplemented")
}

func (o *OrderRepository) FindByID(ctx context.Context, id string) (*domain_orders.Order, error) {
	panic("unimplemented")
}

func (o *OrderRepository) UpdateByID(ctx context.Context, id string, products []string) error {
	panic("unimplemented")
}

func (o *OrderRepository) UpdateStatus(ctx context.Context, id string, status domain_orders.OrderStatus) error {
	panic("unimplemented")
}
