package update_product

import (
	"context"
)

type Input struct {
	ID              string  `uri:"id" swaggerignore:"true"`
	Name            string  `json:"name,omitempty"`
	Category        string  `json:"category,omitempty" enums:"Lanche,Acompanhamento,Bebida,Sobremesa"`
	Price           float64 `json:"price,omitempty"`
	Description     string  `json:"description,omitempty"`
	ImageUrl        string  `json:"imageUrl,omitempty"`
	PreparationTime int64   `json:"preparationTime"`
}

type IUpdateProductUseCase interface {
	Execute(ctx context.Context, input Input) error
}
