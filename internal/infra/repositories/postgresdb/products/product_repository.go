package postgresdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	query := "SELECT id, name, category, price, description, image_url, is_available, created_at, updated_at FROM products WHERE id >= $1 ORDER BY id LIMIT $2;"

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

func (p *ProductRepository) FindByID(ctx context.Context, id int64) (*domain_products.Product, error) {
	query := "SELECT id, name, category, price, description, image_url, is_available, created_at, updated_at FROM products WHERE id = $1 LIMIT 1;"

	row := p.db.QueryRowContext(ctx, query, id)

	var product domain_products.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Category, &product.Price, &product.Description, &product.ImageUrl, &product.IsAvailable, &product.CreatedAt, &product.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}

		return nil, err
	}

	return &product, nil
}

func (p *ProductRepository) Update(ctx context.Context, id int64, product *domain_products.Product) error {
	query := "UPDATE products SET "
	args := []any{}
	setClauses := []string{}
	i := 1

	if product.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", i))
		args = append(args, product.Name)
		i++
	}

	if product.Category != "" {
		setClauses = append(setClauses, fmt.Sprintf("category = $%d", i))
		args = append(args, product.Category)
		i++
	}

	if product.Price > 0 {
		setClauses = append(setClauses, fmt.Sprintf("price = $%d", i))
		args = append(args, product.Price)
		i++
	}

	if product.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", i))
		args = append(args, product.Description)
		i++
	}

	if product.ImageUrl != "" {
		setClauses = append(setClauses, fmt.Sprintf("image_url = $%d", i))
		args = append(args, product.ImageUrl)
		i++
	}

	if len(setClauses) == 0 {
		return errors.New("there are no changes to be made")
	}

	query += fmt.Sprintf("%s WHERE id = $%d;", strings.Join(setClauses, ", "), i)
	args = append(args, id)

	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateAvailability(ctx context.Context, id int64, enable bool) error {
	query := "UPDATE products SET is_available = $1 WHERE id = $2;"

	_, err := p.db.ExecContext(ctx, query, enable, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM products WHERE id = $1;"

	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
