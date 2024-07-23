package repositories

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/models"
)

type IProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	Find(ctx context.Context, offset, limit int64) ([]*models.Product, error)
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	Update(ctx context.Context, id int64, product *models.Product) error
	Delete(ctx context.Context, id int64) error
}

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) error
}
