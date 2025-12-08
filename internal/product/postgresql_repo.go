package product

import (
	"context"
	"database/sql"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) ProductRepository {
	return &postgresRepository{db}
}

func (r *postgresRepository) Save(ctx context.Context, p Product) error {
	query := "INSERT INTO products"
}

func (r *postgresRepository) FindAll(ctx context.Context) ([]Product, error) {
	panic("unimplemented")
}

func (r *postgresRepository) FindByID(ctx context.Context, id string) (Product, error) {
	panic("unimplemented")
}
