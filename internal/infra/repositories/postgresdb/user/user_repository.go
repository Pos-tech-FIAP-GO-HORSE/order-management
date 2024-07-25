package user

import (
	"context"
	"database/sql"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/ports/user/find_user_by_cpf"

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

func (u *UserRepository) FindByCpf(ctx context.Context, cpf int64) (*find_user_by_cpf.User, error) {
	query := "SELECT id, full_name, email, cpf FROM users WHERE cpf = $1 LIMIT 1;"

	row := u.db.QueryRowContext(ctx, query, cpf)

	var user find_user_by_cpf.User

	if err := row.Scan(&user.ID, &user.FullName, user.Email, &user.CPF); err != nil {
		return nil, err
	}

	return &user, nil
}
