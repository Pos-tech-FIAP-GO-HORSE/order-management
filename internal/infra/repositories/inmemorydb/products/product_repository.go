package inmemorydb

import (
	"context"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/core/domain/models"
	"github.com/Pos-tech-FIAP-GO-HORSE/order-management/internal/infra/repositories"
)

type ProductRepository struct {
	products []*models.Product
}

func NewProductRepository() repositories.IProductRepository {
	return &ProductRepository{
		products: make([]*models.Product, 0),
	}
}

func (p *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	p.products = append(p.products, product)
	return nil
}

func (p *ProductRepository) Find(ctx context.Context, offset, limit int64) ([]*models.Product, error) {
	panic("unimplemented")
}

func (p *ProductRepository) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	panic("unimplemented")
}

// TODO: update just the provided values
func (p *ProductRepository) Update(ctx context.Context, id int64, product *models.Product) error {
	panic("unimplemented")
}

func (p *ProductRepository) Delete(ctx context.Context, id int64) error {
	/*for index, value := range p.products {
		if strconv.Itoa(value.ID) == id {
			p.products = append(p.products[:index], p.products[index+1:]...)
		}
	}

	return nil*/
	panic("unimplemented")
}
