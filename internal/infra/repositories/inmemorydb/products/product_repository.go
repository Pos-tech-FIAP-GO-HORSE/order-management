package inmemorydb

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type ProductRepository struct {
	products []*entity.Product
}

func NewProductRepository() repositories.IProductRepository {
	return &ProductRepository{
		products: make([]*entity.Product, 0),
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *entity.Product) error {
	p.products = append(p.products, product)
	return nil
}

func (p *ProductRepository) Find(ctx context.Context, offset, limit int64) ([]*entity.Product, error) {
	panic("unimplemented")
}

func (p *ProductRepository) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	panic("unimplemented")
}

func (p *ProductRepository) FindByCategory(ctx context.Context, category string) ([]*entity.Product, error) {
	panic("unimplemented")
}

func (p *ProductRepository) UpdateByID(ctx context.Context, id string, product *entity.UpdateProduct) error {
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
