package users

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UserRepository struct {
	users []*entity.User
}

func NewUserRepository() repositories.IUserRepository {
	return &UserRepository{
		users: make([]*entity.User, 0),
	}
}

func (u UserRepository) Create(ctx context.Context, user *entity.User) error {
	panic("unimplemented")
}

func (u *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	panic("unimplemented")
}

func (u *UserRepository) FindByCpf(ctx context.Context, cpf string) (*entity.User, error) {
	panic("unimplemented")
}
