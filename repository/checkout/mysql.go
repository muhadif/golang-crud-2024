package checkout

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) repository.CheckoutRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) CreateCart(ctx context.Context, req *entity.CreateCheckoutSession) error {
	return r.db.WithContext(ctx).Create(req).Error
}

func (r repo) CreateCheckout(ctx context.Context, req *entity.CreateCheckoutSession) error {
	return r.db.WithContext(ctx).Table("checkout").Create(req.CartItems).Error
}

func (r repo) GetCurrentCheckout(ctx context.Context, userSerial string) ([]*entity.CartCheckoutItem, error) {
	var items []*entity.CartCheckoutItem

	err := r.db.WithContext(ctx).Table("checkout").Where("user_serial = ?", userSerial).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, err
}

func (r repo) DeleteCheckout(ctx context.Context, userSerial string) error {
	return r.db.WithContext(ctx).Table("checkout").Where("user_serial = ?", userSerial).Error
}
