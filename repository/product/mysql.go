package product

import (
	"context"
	"errors"
	"golang-crud-2024/core/entity"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/fault"
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

func (r repo) GetProductBySerial(ctx context.Context, serial string) (*entity.Product, error) {
	var product *entity.Product
	if err := r.db.WithContext(ctx).Where("serial = ?", serial).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fault.ErrorDictionary(fault.HTTPNotFound, coreErr.ErrProductNotFound)
		}

		return nil, err
	}

	return product, nil
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
