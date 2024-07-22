package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `db:"full_name"`
	Email    string `db:"email"`
	CPF      string `db:"cpf"`
}
