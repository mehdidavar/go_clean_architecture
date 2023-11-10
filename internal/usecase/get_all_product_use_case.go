package usecase

import (
	"context"
	"go-clean-code/internal/entity"
	"go-clean-code/internal/repository"
)

type GetAllProductUseCase struct {
	ProductRepository *repository.ProductRepository
}

func NewGetAllProductUseCase(productRepository *repository.ProductRepository) *GetAllProductUseCase {
	return &GetAllProductUseCase{ProductRepository: productRepository}
}

func (uc *GetAllProductUseCase) Execute(ctx context.Context) ([]entity.Product, error) {
	products := uc.ProductRepository.FindAll(ctx)
	return products, nil
}
