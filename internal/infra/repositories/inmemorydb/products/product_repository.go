package inmemorydb

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/products"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type ProductRepository struct {
	products []*products.Product
}

func NewProductRepository() repositories.IProductRepository {
	return &ProductRepository{
		products: make([]*products.Product, 0),
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *products.Product) error {
	p.products = append(p.products, product)
	return nil
}

func (p *ProductRepository) Find(ctx context.Context, offset, limit int64) ([]*products.Product, error) {
	panic("unimplemented")
}

func (p *ProductRepository) FindByID(ctx context.Context, id string) (*products.Product, error) {
	panic("unimplemented")
}

func (p *ProductRepository) Update(ctx context.Context, id string, product *products.UpdateProduct) error {
	panic("unimplemented")
}

func (p *ProductRepository) UpdateAvailability(ctx context.Context, id string, enable bool) error {
	panic("unimplemented")
}

func (p *ProductRepository) Delete(ctx context.Context, id string) error {
	for index, value := range p.products {
		if value.ID == id {
			p.products = append(p.products[:index], p.products[index+1:]...)
		}
	}

	return nil
}
