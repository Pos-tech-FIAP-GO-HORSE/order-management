package find_user_by_cpf

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	CPF       string    `json:"cpf"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Input struct {
	Cpf string `uri:"cpf" json:"cpf"`
}

type Output struct {
	User User `json:"user"`
}

type IFindUserByCpfUseCase interface {
	Execute(ctx context.Context, input Input) (Output, error)
}
