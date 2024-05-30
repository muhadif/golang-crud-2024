package repository

import (
	"context"
	"golang-crud-2024/core/entity"
)

type CartRepository interface {
	CreateCart(ctx context.Context, req *entity.Cart) error
	GetCartByID(ctx context.Context, req *entity.GetCartByID) (*entity.Cart, error)
	GetCart(ctx context.Context, userSerial string) ([]*entity.Cart, error)
	UpdateCart(ctx context.Context, req *entity.Cart) error
	DeleteCart(ctx context.Context, req *entity.DeleteCart) error
}
