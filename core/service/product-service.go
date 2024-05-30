package module

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
)

type ProductService interface {
	GetProduct(ctx context.Context, req entity.GetProductRequest) ([]*entity.Product, error)
}

type product struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &product{
		productRepo: productRepo,
	}
}

func (p product) GetProduct(ctx context.Context, req entity.GetProductRequest) ([]*entity.Product, error) {
	return p.productRepo.GetProduct(ctx, req)
}
