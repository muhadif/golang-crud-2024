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

func (r repo) GetProduct(ctx context.Context, req entity.GetProductRequest) ([]*entity.Product, error) {
	var products []*entity.Product

	query := r.db.Table("product").Joins("JOIN product_product_category ppc ON product.serial = ppc.product_serial")
	if req.ProductCategorySerial != "" {
		query = query.Where("ppc.product_category_serial = ?", req.ProductCategorySerial)
	}
	err := query.Preload("ProductCategories").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
