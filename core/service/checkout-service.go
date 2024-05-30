package module

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
)

type CheckoutService interface {
	CreateCheckout(ctx context.Context, req *entity.CreateCheckoutSession) error
	GetCurrentCheckout(ctx context.Context, userSerial string) ([]*entity.CartCheckoutItem, error)
}

type checkoutService struct {
	checkoutRepository repository.CheckoutRepository
}

func NewCheckoutService(checkoutRepository repository.CheckoutRepository) CheckoutService {
	return checkoutService{checkoutRepository: checkoutRepository}
}

func (c checkoutService) CreateCheckout(ctx context.Context, req *entity.CreateCheckoutSession) error {
	if err := c.checkoutRepository.DeleteCheckout(ctx, req.UserSerial); err != nil {
		return err
	}

	for _, item := range req.CartItems {
		item.UserSerial = req.UserSerial
	}

	return c.checkoutRepository.CreateCheckout(ctx, req)
}

func (c checkoutService) GetCurrentCheckout(ctx context.Context, userSerial string) ([]*entity.CartCheckoutItem, error) {
	return c.checkoutRepository.GetCurrentCheckout(ctx, userSerial)
}
