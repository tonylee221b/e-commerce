package product

import (
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrInvalidPrice    = errors.New("price cannot be negative")
)

type Money struct {
	amount   float32
	currency string
}

func NewMoney(amount float32) (Money, error) {
	if amount < 0.00 {
		return Money{}, ErrInvalidPrice
	}
	return Money{amount: amount, currency: "KRW"}, nil
}

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price Money  `json:"price"`
}

func NewProduct(id, name string, priceAmount float32) (*Product, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	priceVO, err := NewMoney(priceAmount)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:    id,
		Name:  name,
		Price: priceVO,
	}, nil
}
