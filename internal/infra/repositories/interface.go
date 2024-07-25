package repositories

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/find_user_by_cpf"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/users"
)

type IProductRepository interface {
	Create(ctx context.Context, product *products.Product) error
	Find(ctx context.Context, offset, limit int64) ([]*products.Product, error)
	FindByID(ctx context.Context, id int64) (*products.Product, error)
	Update(ctx context.Context, id int64, product *products.Product) error
	UpdateAvailability(ctx context.Context, id int64, enable bool) error
	Delete(ctx context.Context, id int64) error
}

type IUserRepository interface {
	Create(ctx context.Context, user *users.User) error
	FindByCpf(ctx context.Context, cpf int64) (*find_user_by_cpf.User, error)
}
