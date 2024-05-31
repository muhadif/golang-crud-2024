package repository

import (
	"context"
	"golang-crud-2024/core/entity"
)

type ProductCategoryRepository interface {
	GetProductCategory(ctx context.Context) ([]*entity.ProductCategory, error)
}

type ProductCategoryRedisRepository interface {
	GetProductCategory(ctx context.Context) ([]*entity.ProductCategory, error)
	SetProductCategory(ctx context.Context, req []*entity.ProductCategory) error
}
