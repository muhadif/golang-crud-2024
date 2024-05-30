package repository

import (
	"context"
	"golang-crud-2024/core/entity"
)

type CheckoutRepository interface {
	CreateCheckout(ctx context.Context, req *entity.CreateCheckoutSession) error
	GetCurrentCheckout(ctx context.Context, userSerial string) ([]*entity.CartCheckoutItem, error)
	DeleteCheckout(ctx context.Context, userSerial string) error
}
