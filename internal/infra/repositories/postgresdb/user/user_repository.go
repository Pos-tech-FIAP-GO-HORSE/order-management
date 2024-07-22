package user

import (
	"context"
	"database/sql"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/users"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) Create(ctx context.Context, user *users.User) error {
	query := "INSERT INTO users (full_name, email, cpf) VALUES ($1, $2, $3)"

	_, err := u.db.ExecContext(ctx, query, user.FullName, user.Email, user.CPF)
	if err != nil {
		return err
	}

	return nil
}
