package product_category

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	"gorm.io/gorm"
)

func NewProductCategoryRepository(db *gorm.DB) repository.ProductCategoryRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) GetProductCategory(ctx context.Context) ([]*entity.ProductCategory, error) {
	var productCategories []*entity.ProductCategory

	err := r.db.Table("product_category").Find(&productCategories).Error
	if err != nil {
		return nil, err
	}

	return productCategories, nil
}
