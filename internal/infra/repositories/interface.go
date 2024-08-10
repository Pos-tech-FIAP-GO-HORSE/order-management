package repositories

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/payment"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/orders"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/users"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/utils"
)

type IProductRepository interface {
	Create(ctx context.Context, product *products.Product) error
	Find(ctx context.Context, offset, limit int64) ([]*products.Product, error)
	FindByID(ctx context.Context, id string) (*products.Product, error)
	UpdateByID(ctx context.Context, id string, product *products.UpdateProduct) error
	UpdateAvailability(ctx context.Context, id string, enable bool) error
	Delete(ctx context.Context, id string) error
}

type IUserRepository interface {
	Create(ctx context.Context, user *users.User) error
	FindByID(ctx context.Context, id string) (*users.User, error)
	FindByCpf(ctx context.Context, cpf string) (*users.User, error)
}

type IOrderRepository interface {
	Create(ctx context.Context, order *orders.Order) error
	Find(ctx context.Context, filter utils.OrderFilters, offset, limit int64) ([]*orders.Order, error)
	FindByID(ctx context.Context, id string) (*orders.Order, error)
	UpdateByID(ctx context.Context, id string, products []string) error
	UpdateStatus(ctx context.Context, id string, status orders.OrderStatus) error
}

type IPaymentRepository interface {
	Save(ctx context.Context, payment *payment.Payment) error
	GetByID(ctx context.Context, id string) (*payment.Payment, error)
	Update(ctx context.Context, payment *payment.Payment) error
}
