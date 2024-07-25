package user

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/find_user_by_cpf"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type FindUserByCpfUseCase struct {
	UserRepository repositories.IUserRepository
}

func NewFindUserByCpfUseCase(userRepository repositories.IUserRepository) find_user_by_cpf.IFindUserByCpfUseCase {
	return &FindUserByCpfUseCase{
		UserRepository: userRepository,
	}
}

func (f *FindUserByCpfUseCase) Execute(ctx context.Context, input find_user_by_cpf.Input) (find_user_by_cpf.Output, error) {
	user, err := f.UserRepository.FindByCpf(ctx, input.Cpf)
	if err != nil {
		return find_user_by_cpf.Output{}, err
	}

	output := find_user_by_cpf.Output{
		User: find_user_by_cpf.User{
			ID:        user.ID,
			FullName:  user.FullName,
			Email:     user.Email,
			CPF:       user.CPF,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return output, nil
}
