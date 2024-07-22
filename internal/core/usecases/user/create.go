package user

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/users"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/create_user"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type CreateUserUseCase struct {
	UserRepository repositories.IUserRepository
}

func NewCreateUserUseCase(userRepository repositories.IUserRepository) create_user.ICreateUserUseCase {
	return &CreateUserUseCase{UserRepository: userRepository}
}

func (c *CreateUserUseCase) Execute(ctx context.Context, newUser create_user.NewUser) error {
	user, err := users.NewUser(&newUser)
	if err != nil {
		return err
	}

	if err = c.UserRepository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
