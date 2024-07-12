package postgresdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	domain_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repositories.IProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *domain_products.Product) error {
	query := "INSERT INTO products (name, category, price, description, image_url, is_available) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := p.db.ExecContext(ctx, query, product.Name, product.Category, product.Price, product.Description, product.ImageUrl, product.IsAvailable)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Find(ctx context.Context, offset, limit int64) ([]*domain_products.Product, error) {
	query := "SELECT id, name, category, price, description, image_url, is_available, created_at, updated_at FROM products WHERE id >= $1 LIMIT $2;"

	rows, err := p.db.QueryContext(ctx, query, offset, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("products not found")
		}

		return nil, err
	}
	defer rows.Close()

	products := make([]*domain_products.Product, 0)

	for rows.Next() {
		var product domain_products.Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description, &product.ImageUrl, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func (p *ProductRepository) FindByID(ctx context.Context, id int64) (*products.Product, error) {
	query := "SELECT id, name, category, price, description, image_url, is_available, created_at, updated_at FROM products WHERE id = $1 LIMIT 1;"

	row := p.db.QueryRowContext(ctx, query, id)

	var product products.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description, &product.ImageUrl, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}

		return nil, err
	}

	return &product, nil
}

// TODO: update just the provided values
func (p *ProductRepository) Update(ctx context.Context, id int64, product *products.Product) error {
	panic("unimplemented")
}

func (p *ProductRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM products WHERE id = $1;"

	result, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows <= 0 {
		return errors.New("product not found to be deleted")
	}

	return nil
}
