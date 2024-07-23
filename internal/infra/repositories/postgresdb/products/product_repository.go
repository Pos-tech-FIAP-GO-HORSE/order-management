package postgresdb

import (
	"context"
	"errors"
	domain_products "github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/models"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/db/db_gorm"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type ProductRepository struct {
}

func NewProductRepository() repositories.IProductRepository {
	return &ProductRepository{}
}

func (p *ProductRepository) Create(ctx context.Context, product *domain_products.Product) error {

	err := db_gorm.DB.Create(&product)
	if err != nil {
		return errors.New("products not found")
	}

	return nil
}

func (p *ProductRepository) Find(ctx context.Context, offset, limit int64) ([]*domain_products.Product, error) {

	var products []*domain_products.Product
	err := db_gorm.DB.Find(&products)
	if err != nil {
		return nil, errors.New("products not found")
	}

	return products, nil
}

func (p *ProductRepository) FindByID(ctx context.Context, id int64) (*domain_products.Product, error) {

	var product *domain_products.Product
	err := db_gorm.DB.First(&product, id)
	if err != nil {
		return nil, errors.New("products not found")
	}

	return product, nil
}

// TODO: update just the provided values
func (p *ProductRepository) Update(ctx context.Context, id int64, product *domain_products.Product) error {
	panic("unimplemented")
}

func (p *ProductRepository) Delete(ctx context.Context, id int64) error {

	var product *domain_products.Product

	err := db_gorm.DB.First(&product, id)
	if err != nil {
		return errors.New("products not found")
	}
	if product.ID == 0 {
		return errors.New("products not found")
	}

	return nil
}
