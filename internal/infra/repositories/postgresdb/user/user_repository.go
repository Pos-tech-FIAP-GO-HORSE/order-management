package user

import (
	"context"
	"errors"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/models"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/db/db_gorm"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UserRepository struct {
}

func NewUserRepository() repositories.IUserRepository {
	return &UserRepository{}
}

func (u UserRepository) Create(ctx context.Context, user *models.User) error {

	err := db_gorm.DB.Create(&user)

	if err != nil {
		return errors.New("products not found")
	}

	return nil
}
