package orders

import (
	"context"

	domain_orders "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (o *OrderRepository) Find(ctx context.Context, filter utils.OrderFilters, offset int64, limit int64) ([]*domain_orders.Order, error) {
	filters := bson.M{}

	if filter.Status != "" {
		filters["status"] = filter.Status
	}

	cursor, err := o.collection.Find(ctx, filters, &options.FindOptions{
		Skip:  &offset,
		Limit: &limit,
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	orders := make([]*domain_orders.Order, 0)

	for cursor.Next(ctx) {
		var order domain_orders.Order
		if err = cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
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
