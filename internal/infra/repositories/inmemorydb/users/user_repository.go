package users

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/users"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UserRepository struct {
	users []*users.User
}

func NewUserRepository() repositories.IUserRepository {
	return &UserRepository{
		users: make([]*users.User, 0),
	}
}

func (u UserRepository) Create(ctx context.Context, user *users.User) error {
	panic("unimplemented")
}

func (u *UserRepository) FindByCpf(ctx context.Context, cpf string) (*users.User, error) {
	panic("unimplemented")
}
