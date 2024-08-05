package orders

import (
	"errors"
	"time"
)

type OrderStatus string

const (
	Criado       OrderStatus = "Criado"
	EmPreparacao OrderStatus = "Em preparação"
	Pronto       OrderStatus = "Pronto"
	Finalizado   OrderStatus = "Finalizado"
)

type Order struct {
	ID         string      `bson:"_id,omitempty" json:"id"`
	UserID     string      `bson:"userId" json:"userId"`
	Items      []*Item     `bson:"items" json:"items"`
	TotalPrice float64     `bson:"totalPrice" json:"totalPrice"`
	Status     OrderStatus `bson:"status" json:"status"`
	CreatedAt  time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time   `bson:"updatedAt" json:"updatedAt"`
}

type Item struct {
	ID       string  `bson:"id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Comments string  `bson:"comments" json:"comments"`
}

func NewOrder(userID string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("items could not be empty")
	}

	totalPrice := calculateTotalPrice(items)
	if totalPrice <= 0 {
		return nil, errors.New("total price not provided")
	}

	return &Order{
		UserID:     userID,
		Items:      items,
		TotalPrice: totalPrice,
		Status:     Criado,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func NewItem(id, name, comments string, price float64) (*Item, error) {
	if id == "" {
		return nil, errors.New("item id not provided")
	}

	if name == "" {
		return nil, errors.New("name not provided")
	}

	if price <= 0 {
		return nil, errors.New("item price not provided")
	}

	return &Item{
		ID:       id,
		Name:     name,
		Price:    price,
		Comments: comments,
	}, nil
}

func calculateTotalPrice(items []*Item) float64 {
	var totalPrice float64
	for _, item := range items {
		totalPrice += item.Price
	}

	return totalPrice
}
