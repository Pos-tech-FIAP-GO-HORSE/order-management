package repositories

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
	valueobjects "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/valueObjects"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/utils"
)

type IProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	Find(ctx context.Context, offset, limit int64) ([]*entity.Product, error)
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	FindByCategory(ctx context.Context, category string) ([]*entity.Product, error)
	UpdateByID(ctx context.Context, id string, product *entity.UpdateProduct) error
	UpdateAvailability(ctx context.Context, id string, enable bool) error
	Delete(ctx context.Context, id string) error
}

type IUserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindByCpf(ctx context.Context, cpf string) (*entity.User, error)
}

type IOrderRepository interface {
	Create(ctx context.Context, order *entity.Order) (entity.OrderOutput, error)
	Find(ctx context.Context, filter utils.OrderFilters, offset, limit int64) ([]*entity.Order, error)
	FindByID(ctx context.Context, id string) (*entity.Order, error)
	UpdateByID(ctx context.Context, id string, order *entity.UpdateOrder) error
	UpdateStatusByID(ctx context.Context, id, status valueobjects.OrderStatusType) error
}
