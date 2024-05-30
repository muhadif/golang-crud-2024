package module

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
)

type CartService interface {
	GetCart(ctx context.Context, userSerial string) ([]*entity.Cart, error)
	GetCartByID(ctx context.Context, req *entity.GetCartByID) (*entity.Cart, error)
	CreateCart(ctx context.Context, req *entity.CreateCart) error
	UpdateCart(ctx context.Context, req *entity.UpdateCart) error
	DeleteCart(ctx context.Context, req *entity.DeleteCart) error
}

type cartService struct {
	CartRepo repository.CartRepository
}

func NewCartService(cartRepository repository.CartRepository) CartService {
	return cartService{CartRepo: cartRepository}
}

func (c cartService) GetCart(ctx context.Context, userSerial string) ([]*entity.Cart, error) {
	return c.CartRepo.GetCart(ctx, userSerial)
}

func (c cartService) GetCartByID(ctx context.Context, req *entity.GetCartByID) (*entity.Cart, error) {
	return c.CartRepo.GetCartByID(ctx, req)
}

func (c cartService) CreateCart(ctx context.Context, req *entity.CreateCart) error {
	return c.CartRepo.CreateCart(ctx, &entity.Cart{
		UserSerial:    req.UserSerial,
		ProductSerial: req.ProductSerial,
		Quantity:      req.Quantity,
	})
}

func (c cartService) UpdateCart(ctx context.Context, req *entity.UpdateCart) error {
	if req.Quantity <= 0 {
		return c.CartRepo.DeleteCart(ctx, &entity.DeleteCart{
			ID:         req.ID,
			UserSerial: req.UserSerial,
		})
	}

	return c.CartRepo.UpdateCart(ctx, &entity.Cart{
		ID:            req.ID,
		UserSerial:    req.UserSerial,
		ProductSerial: req.ProductSerial,
		Quantity:      req.Quantity,
	})
}

func (c cartService) DeleteCart(ctx context.Context, req *entity.DeleteCart) error {
	return c.CartRepo.DeleteCart(ctx, &entity.DeleteCart{
		UserSerial: req.UserSerial,
		ID:         req.ID,
	})
}
