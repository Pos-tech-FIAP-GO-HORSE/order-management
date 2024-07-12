package products

import (
	"errors"
	"time"
)

type Product struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Category    string    `db:"category"`
	Price       float64   `db:"price"`
	Description string    `db:"description"`
	ImageUrl    string    `db:"image_url"`
	IsAvailable bool      `db:"is_available"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewProduct(name, category, description, imageUrl string, price float64) (*Product, error) {
	if name == "" {
		return nil, errors.New("name not provided")
	}

	if category == "" {
		return nil, errors.New("category not provided")
	}

	if description == "" {
		return nil, errors.New("description not provided")
	}

	if imageUrl == "" {
		return nil, errors.New("image url not provided")
	}

	if price <= 0 {
		return nil, errors.New("price not provided")
	}

	return &Product{
		Name:        name,
		Category:    category,
		Price:       price,
		Description: description,
		ImageUrl:    imageUrl,
		IsAvailable: true,
	}, nil
}
