package orders

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	domain_orders "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
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

func (o *OrderRepository) Create(ctx context.Context, order *domain_orders.Order) (domain_orders.OrderOutput, error) {
	result, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return domain_orders.OrderOutput{}, err
	}

	orderID := result.InsertedID.(primitive.ObjectID)

	return domain_orders.OrderOutput{ID: orderID.Hex()}, nil
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := o.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err := result.Err(); err != nil {
		return nil, err
	}

	var order domain_orders.Order
	if err = result.Decode(&order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *OrderRepository) UpdateByID(ctx context.Context, id string, order *domain_orders.UpdateOrder) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := o.collection.UpdateByID(ctx, objectID, bson.M{"$set": order})
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no updates have been made")
	}

	_, err = o.collection.UpdateByID(ctx, objectID, bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) UpdateStatusByID(ctx context.Context, id string, status string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := o.collection.UpdateByID(ctx, objectID, bson.M{
		"$set": bson.M{
			"status": status,
		},
	})

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no updates have been made")
	}

	_, err = o.collection.UpdateByID(ctx, objectID, bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
