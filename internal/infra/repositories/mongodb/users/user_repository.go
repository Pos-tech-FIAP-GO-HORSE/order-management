package users

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/users"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) repositories.IUserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (u *UserRepository) Create(ctx context.Context, user *users.User) error {
	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindByCpf(ctx context.Context, cpf string) (*users.User, error) {

	result := u.collection.FindOne(ctx, bson.M{"cpf": cpf})

	var user users.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
