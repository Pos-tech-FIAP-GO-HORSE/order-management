package entity

import (
	"errors"
	"time"
)

type User struct {
	ID        string    `bson:"_id,omitempty" db:"id"`
	FullName  string    `bson:"fullName,omitempty" db:"full_name"`
	Email     string    `bson:"email,omitempty" db:"email"`
	CPF       string    `bson:"cpf,omitempty" db:"cpf"`
	CreatedAt time.Time `bson:"createdAt,omitempty" db:"created_at"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" db:"updated_at"`
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
		FullName:  fullName,
		Email:     email,
		CPF:       cpf,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
