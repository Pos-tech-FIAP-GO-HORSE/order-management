package create_user

import "context"

type NewUser struct {
	FullName string `json:"fullName" validate:"min=3,max=100"`
	CPF      string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	Email    string `json:"email"`
}

type ICreateUserUseCase interface {
	Execute(ctx context.Context, createUser NewUser) error
}
