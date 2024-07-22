package users

import (
	"errors"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/create_user"
	"time"
)

type User struct {
	ID        int64     `db:"id"`
	FullName  string    `db:"full_name"`
	Email     string    `db:"email"`
	CPF       string    `db:"cpf"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUser(newUser *create_user.NewUser) (*User, error) {
	if newUser.FullName == "" {
		return nil, errors.New("name not provided")
	}
	if newUser.Email == "" {
		return nil, errors.New("email not provided")
	}
	if newUser.CPF == "" {
		return nil, errors.New("email not provided")
	}

	return &User{
		FullName: newUser.FullName,
		Email:    newUser.Email,
		CPF:      newUser.CPF,
	}, nil
}
