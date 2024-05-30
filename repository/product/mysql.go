package product

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) repository.ProductRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) GetProduct(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product

	err := r.db.Table("product").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
