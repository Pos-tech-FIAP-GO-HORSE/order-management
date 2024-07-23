package users

import (
	"errors"
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

func NewUser(fullName string, email string, cpf string) (*User, error) {
	if fullName == "" {
		return nil, errors.New("name not provided")
	}
	if email == "" {
		return nil, errors.New("email not provided")
	}
	if cpf == "" {
		return nil, errors.New("email not provided")
	}

	return &User{
		FullName: fullName,
		Email:    email,
		CPF:      cpf,
	}, nil
}
