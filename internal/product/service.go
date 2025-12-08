package product

import (
	"context"
	"time"
)

type ProductService struct {
	repo ProductRepository
}

func NewService(repo ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, id, name string, priceAmount float32) error {
	p, err := NewProduct(id, name, priceAmount)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return s.repo.Save(ctx, *p)
}

func (s *ProductService) ListProducts(ctx context.Context) ([]Product, error) {
	return s.repo.FindAll(ctx)
}
