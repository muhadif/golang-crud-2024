package module

import (
	"context"
	"golang-crud-2024/core/entity"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/fault"
)

type CheckoutService interface {
	CreateCheckout(ctx context.Context, req *entity.CreateCheckoutSession) error
	GetCurrentCheckout(ctx context.Context, userSerial string) (*entity.GetCheckoutSessionResponse, error)
	DeleteCheckout(ctx context.Context, userSerial string) error
}

type checkoutService struct {
	checkoutRepository repository.CheckoutRepository
	cartRepo           repository.CartRepository
}

func NewCheckoutService(checkoutRepository repository.CheckoutRepository, cartRepository repository.CartRepository) CheckoutService {
	return checkoutService{checkoutRepository: checkoutRepository, cartRepo: cartRepository}
}

func (c checkoutService) CreateCheckout(ctx context.Context, req *entity.CreateCheckoutSession) error {
	if err := c.checkoutRepository.DeleteCheckout(ctx, req.UserSerial); err != nil {
		return err
	}

	for _, item := range req.CartItems {
		cart, err := c.cartRepo.GetCartByID(ctx, &entity.GetCartByID{
			ID:         item.CartID,
			UserSerial: req.UserSerial,
		})
		if err != nil {
			return err
		}

		if cart == nil {
			return fault.ErrorDictionary(fault.HTTPNotFound, coreErr.ErrCartNotFOund)
		}

		item.UserSerial = req.UserSerial
	}

	return c.checkoutRepository.CreateCheckout(ctx, req)
}

func (c checkoutService) GetCurrentCheckout(ctx context.Context, userSerial string) (*entity.GetCheckoutSessionResponse, error) {
	currentCart, err := c.checkoutRepository.GetCurrentCheckout(ctx, userSerial)
	if err != nil {
		return nil, err
	}

	var totalPrice float64
	for _, item := range currentCart {
		totalPrice += float64(item.Quantity) * item.Product.Price
	}

	return &entity.GetCheckoutSessionResponse{
		CartItems: currentCart,
		Total:     totalPrice,
	}, nil
}

func (c checkoutService) DeleteCheckout(ctx context.Context, userSerial string) error {
	return c.checkoutRepository.DeleteCheckout(ctx, userSerial)
}
