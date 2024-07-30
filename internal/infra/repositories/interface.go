package repositories

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/users"
)

type IProductRepository interface {
	Create(ctx context.Context, product *products.Product) error
	Find(ctx context.Context, offset, limit int64) ([]*products.Product, error)
	FindByID(ctx context.Context, id string) (*products.Product, error)
	FindByCategory(ctx context.Context, category string) (*products.Product, error)
	Update(ctx context.Context, id string, product *products.UpdateProduct) error
	UpdateAvailability(ctx context.Context, id string, enable bool) error
	Delete(ctx context.Context, id string) error
}

type IUserRepository interface {
	Create(ctx context.Context, user *users.User) error
	FindByCpf(ctx context.Context, cpf string) (*users.User, error)
}
