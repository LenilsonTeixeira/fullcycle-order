package entity

import (
	"errors"

	"github.com/google/uuid"
)

type OrderRepositoryInterface interface {
	Save(order *Order) error
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    uuid.New().String(),
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		panic(err)
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price == 0 {
		return errors.New("invalid price")
	}
	if o.Tax == 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price * o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
