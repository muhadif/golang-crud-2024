package module

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
)

type ProductCategoryService interface {
	GetProductCategory(ctx context.Context) ([]*entity.ProductCategory, error)
}

type productCategory struct {
	productCategoryRepo repository.ProductCategoryRepository
}

func NewProductCategoryService(productCategoryRepo repository.ProductCategoryRepository) ProductCategoryService {
	return &productCategory{
		productCategoryRepo: productCategoryRepo,
	}
}

func (p productCategory) GetProductCategory(ctx context.Context) ([]*entity.ProductCategory, error) {
	return p.productCategoryRepo.GetProductCategory(ctx)
}
