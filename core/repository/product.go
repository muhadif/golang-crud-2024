package repository

import (
	"context"
	"golang-crud-2024/core/entity"
)

type ProductRepository interface {
	GetProduct(ctx context.Context) ([]*entity.Product, error)
}
