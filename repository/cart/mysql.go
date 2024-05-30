package cart

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	"gorm.io/gorm"
	"time"
)

func NewRepository(db *gorm.DB) repository.CartRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) CreateCart(ctx context.Context, req *entity.Cart) error {
	return r.db.WithContext(ctx).Create(req).Error
}

func (r repo) GetCartByID(ctx context.Context, req *entity.GetCartByID) (*entity.Cart, error) {
	var cart *entity.Cart
	err := r.db.WithContext(ctx).
		Where("id = ?", req.ID).
		Where("user_serial = ?", req.UserSerial).
		Where("deleted_at IS NULL").
		Preload("Product").
		First(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r repo) GetCart(ctx context.Context, userSerial string) ([]*entity.Cart, error) {
	var carts []*entity.Cart
	err := r.db.WithContext(ctx).
		Where("user_serial = ?", userSerial).
		Where("deleted_at IS NULL").
		Preload("Product").
		Find(&carts).Error
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (r repo) UpdateCart(ctx context.Context, req *entity.Cart) error {
	return r.db.Where("id = ?", req.ID).
		Where("user_serial = ?", req.UserSerial).Updates(&req).Error
}

func (r repo) DeleteCart(ctx context.Context, req *entity.DeleteCart) error {
	return r.db.Model(&entity.Cart{}).
		Where("id = ?", req.ID).
		Where("user_serial = ?", req.UserSerial).Update("deleted_at", time.Now()).Error
}
