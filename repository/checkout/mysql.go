package checkout

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	"gorm.io/gorm"
	"time"
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

func (r repo) GetCurrentCheckout(ctx context.Context, userSerial string) ([]*entity.Cart, error) {
	var items []*entity.Cart

	err := r.db.WithContext(ctx).Table("checkout").Preload("Product").
		Joins("JOIN cart ON cart.id = checkout.cart_id").
		Where("checkout.user_serial = ?", userSerial).
		Where("checkout.deleted_at IS NULL").
		Select("cart.*").
		Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, err
}

func (r repo) DeleteCheckout(ctx context.Context, userSerial string) error {
	return r.db.WithContext(ctx).Table("checkout").Where("user_serial = ?", userSerial).Update("deleted_at", time.Now()).Error
}
