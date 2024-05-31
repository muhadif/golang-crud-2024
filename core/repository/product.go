package repository

import (
	"context"
	"golang-crud-2024/core/entity"
)

type ProductRepository interface {
	GetProduct(ctx context.Context, req entity.GetProductRequest) ([]*entity.Product, error)
	GetProductBySerial(ctx context.Context, serial string) (*entity.Product, error)
	RollbackStock(ctx context.Context, e *entity.RollbackStockRequest) error
}
