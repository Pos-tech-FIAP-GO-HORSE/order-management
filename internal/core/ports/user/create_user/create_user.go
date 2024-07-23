package create_user

import "context"

type Input struct {
	FullName string `json:"fullName"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
}

type ICreateUserUseCase interface {
	Execute(ctx context.Context, createUser Input) error
}
