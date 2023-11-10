package usecase

import (
	"context"
	"go-clean-code/internal/entity"
	"go-clean-code/internal/repository"
)

type InsertProductUseCase struct {
	ProductRepository *repository.ProductRepository
}

func NewInsertProductUseCase(productRepository *repository.ProductRepository) *InsertProductUseCase {
	return &InsertProductUseCase{ProductRepository: productRepository}
}

func (insertUseCase *InsertProductUseCase) Execute(ctx context.Context, product entity.Product) (entity.Product, error) {
	insertedProduct := insertUseCase.ProductRepository.Insert(ctx, product)
	return insertedProduct, nil
}
