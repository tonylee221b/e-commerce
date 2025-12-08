package product

import "context"

type ProductRepository interface {
	Save(ctx context.Context, p Product) error
	FindAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id string) (Product, error)
}
