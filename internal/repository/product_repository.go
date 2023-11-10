package repository

import (
	"context"
	"database/sql"
	"go-clean-code/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(dB *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: dB,
	}
}

func (repo *ProductRepository) Insert(ctx context.Context, product entity.Product) entity.Product {

	return product
}

func (repo *ProductRepository) FindAll(ctx context.Context) []entity.Product {

	products := []entity.Product{
		{ID: 1, Name: "Product1"},
		{ID: 2, Name: "Product2"},
	}
	return products
}
