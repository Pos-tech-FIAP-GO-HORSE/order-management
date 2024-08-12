package orders

import (
	"errors"
	"time"
)

type OrderStatus string

const (
	Received        OrderStatus = "Received"
	AwaitingPayment OrderStatus = "Awaiting Payment"
	Confirmed       OrderStatus = "Confirmed"
	Preparing       OrderStatus = "Preparing"
	Ready           OrderStatus = "Ready"
	Finished        OrderStatus = "Finished"
	Canceled        OrderStatus = "Canceled"
)

type Order struct {
	ID                       string      `bson:"_id,omitempty" json:"id"`
	UserID                   string      `bson:"userId" json:"userId"`
	Items                    []*Item     `bson:"items" json:"items"`
	TotalPrice               float64     `bson:"totalPrice" json:"totalPrice"`
	Status                   OrderStatus `bson:"status" json:"status"`
	EstimatedPreparationTime int64       `bson:"estimatedPreparationTime" json:"estimatedPreparationTime"`
	CreatedAt                time.Time   `bson:"createdAt" json:"createdAt"`
	UpdatedAt                time.Time   `bson:"updatedAt" json:"updatedAt"`
}

type UpdateOrder struct {
	UserID                   string      `bson:"userId,omitempty" json:"userId"`
	Items                    []*Item     `bson:"items,omitempty" json:"items"`
	TotalPrice               float64     `bson:"totalPrice,omitempty" json:"totalPrice"`
	Status                   OrderStatus `bson:"status,omitempty" json:"status"`
	EstimatedPreparationTime int64       `bson:"estimatedPreparationTime,omitempty" json:"estimatedPreparationTime"`
}

type Item struct {
	ID              string  `bson:"id" json:"id"`
	Name            string  `bson:"name" json:"name"`
	Price           float64 `bson:"price" json:"price"`
	Comments        string  `bson:"comments" json:"comments"`
	PreparationTime int64   `bson:"-" json:"-"`
}

func NewOrder(userID string, items []*Item) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("items could not be empty")
	}

	return &Order{
		UserID:    userID,
		Items:     items,
		Status:    Received,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func NewItem(id, name, comments string, price float64, preparationTime int64) (*Item, error) {
	if id == "" {
		return nil, errors.New("item id not provided")
	}

	if name == "" {
		return nil, errors.New("name not provided")
	}

	if price <= 0 {
		return nil, errors.New("item price not provided")
	}

	if preparationTime <= 0 {
		return nil, errors.New("item preparation time not provided")
	}

	return &Item{
		ID:              id,
		Name:            name,
		Price:           price,
		Comments:        comments,
		PreparationTime: preparationTime,
	}, nil
}

func (o *Order) CalculateTotalPrice() {
	if len(o.Items) > 0 {
		var totalPrice float64
		for _, item := range o.Items {
			totalPrice += item.Price
		}

		o.TotalPrice = totalPrice
	}
}

func (o *Order) CalculateEstimatedPreparationTime() {
	if len(o.Items) > 0 {
		var preparationTime int64
		for _, item := range o.Items {
			preparationTime += item.PreparationTime
		}

		o.EstimatedPreparationTime = preparationTime / int64(len(o.Items))
	}
}
