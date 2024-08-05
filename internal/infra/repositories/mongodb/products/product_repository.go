package products

import (
	"context"
	"errors"
	"time"

	domain_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(collection *mongo.Collection) repositories.IProductRepository {
	return &ProductRepository{
		collection: collection,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *domain_products.Product) error {
	_, err := p.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Find(ctx context.Context, offset, limit int64) ([]*domain_products.Product, error) {
	cursor, err := p.collection.Find(ctx, bson.M{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	products := make([]*domain_products.Product, 0)

	for cursor.Next(ctx) {
		var product domain_products.Product
		if err = cursor.Decode(&product); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *ProductRepository) FindByID(ctx context.Context, id string) (*domain_products.Product, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := p.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err = result.Err(); err != nil {
		return nil, err
	}

	var product domain_products.Product
	if err := result.Decode(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepository) UpdateByID(ctx context.Context, id string, product *domain_products.UpdateProduct) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := p.collection.UpdateByID(ctx, objectID, bson.M{"$set": product})
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no updates have been made")
	}

	_, err = p.collection.UpdateByID(ctx, objectID, bson.M{
		"$set": bson.M{
			"updatedAt": time.Now(),
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateAvailability(ctx context.Context, id string, enable bool) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := p.collection.UpdateByID(ctx, objectID, bson.M{
		"$set": bson.M{
			"isAvailable": enable,
			"updatedAt":   time.Now(),
		},
	})

	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no updates have been made")
	}

	return nil
}

func (p *ProductRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := p.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no deletions have been made")
	}

	return nil
}
