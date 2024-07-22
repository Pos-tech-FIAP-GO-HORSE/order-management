package create_user

import "context"

type NewUser struct {
	FullName string `json:"fullName"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
}

type ICreateUserUseCase interface {
	Execute(ctx context.Context, createUser NewUser) error
}
