package products

import (
	"errors"
	"time"
)

type Product struct {
	ID              string    `bson:"_id,omitempty" json:"id" db:"id"`
	Name            string    `bson:"name,omitempty" json:"name" db:"name"`
	Category        string    `bson:"category,omitempty" json:"category" db:"category"`
	Price           float64   `bson:"price,omitempty" json:"price" db:"price"`
	Description     string    `bson:"description,omitempty" json:"description" db:"description"`
	ImageUrl        string    `bson:"imageUrl,omitempty" json:"imageUrl" db:"image_url"`
	IsAvailable     bool      `bson:"isAvailable,omitempty" json:"isAvailable" db:"is_available"`
	PreparationTime int64     `bson:"preparationTime,omitempty" json:"preparationTime" db:"preparation_time"`
	CreatedAt       time.Time `bson:"createdAt,omitempty" json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `bson:"updatedAt,omitempty" json:"updatedAt" db:"updated_at"`
}

type UpdateProduct struct {
	Name            string  `bson:"name,omitempty" json:"name" db:"name"`
	Category        string  `bson:"category,omitempty" json:"category" db:"category"`
	Price           float64 `bson:"price,omitempty" json:"price" db:"price"`
	Description     string  `bson:"description,omitempty" json:"description" db:"description"`
	ImageUrl        string  `bson:"imageUrl,omitempty" json:"imageUrl" db:"image_url"`
	PreparationTime int64   `bson:"preparationTime,omitempty" json:"preparationTime" db:"preparation_time"`
}

func NewProduct(name, category, description, imageUrl string, price float64, preparationTime int64) (*Product, error) {
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

	if preparationTime <= 0 {
		return nil, errors.New("preparation time not provided")
	}

	return &Product{
		Name:            name,
		Category:        category,
		Price:           price,
		Description:     description,
		ImageUrl:        imageUrl,
		IsAvailable:     true,
		PreparationTime: preparationTime,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}
