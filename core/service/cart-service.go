package module

import (
	"context"
	"golang-crud-2024/core/entity"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/fault"
)

type CartService interface {
	GetCart(ctx context.Context, userSerial string) ([]*entity.Cart, error)
	GetCartByID(ctx context.Context, req *entity.GetCartByID) (*entity.Cart, error)
	CreateCart(ctx context.Context, req *entity.CreateCart) error
	UpdateCart(ctx context.Context, req *entity.UpdateCart) error
	DeleteCart(ctx context.Context, req *entity.DeleteCart) error
}

type cartService struct {
	CartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartService(cartRepository repository.CartRepository, productRepo repository.ProductRepository) CartService {
	return cartService{CartRepo: cartRepository, productRepo: productRepo}
}

func (c cartService) GetCart(ctx context.Context, userSerial string) ([]*entity.Cart, error) {
	return c.CartRepo.GetCart(ctx, userSerial)
}

func (c cartService) GetCartByID(ctx context.Context, req *entity.GetCartByID) (*entity.Cart, error) {
	return c.CartRepo.GetCartByID(ctx, req)
}

func (c cartService) CreateCart(ctx context.Context, req *entity.CreateCart) error {
	product, err := c.productRepo.GetProductBySerial(ctx, req.ProductSerial)
	if err != nil {
		return err
	}

	quantity := req.Quantity
	existingCart, _ := c.CartRepo.GetCartsByUserSerialAndProductSerial(ctx, req.UserSerial, req.ProductSerial)
	if existingCart != nil {
		existingCart.Quantity += req.Quantity
		quantity += existingCart.Quantity
	}

	if product.Stock < quantity {
		return fault.ErrorDictionary(fault.HTTPBadRequestError, coreErr.ErrProductStock)
	}

	if existingCart != nil {
		return c.CartRepo.UpdateCart(ctx, existingCart)
	}

	return c.CartRepo.CreateCart(ctx, &entity.Cart{
		UserSerial:    req.UserSerial,
		ProductSerial: req.ProductSerial,
		Quantity:      req.Quantity,
	})
}

func (c cartService) UpdateCart(ctx context.Context, req *entity.UpdateCart) error {
	cart, err := c.CartRepo.GetCartByID(ctx, &entity.GetCartByID{
		ID:         req.ID,
		UserSerial: req.UserSerial,
	})
	if err != nil {
		return err
	}

	if cart == nil {
		return fault.ErrorDictionary(fault.HTTPNotFound, coreErr.ErrCartNotFOund)
	}

	if req.Quantity > cart.Product.Stock {
		return fault.ErrorDictionary(fault.HTTPBadRequestError, coreErr.ErrProductStock)
	}

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
	cart, err := c.CartRepo.GetCartByID(ctx, &entity.GetCartByID{
		ID:         req.ID,
		UserSerial: req.UserSerial,
	})
	if err != nil {
		return err
	}

	if cart == nil {
		return fault.ErrorDictionary(fault.HTTPNotFound, coreErr.ErrCartNotFOund)
	}

	return c.CartRepo.DeleteCart(ctx, &entity.DeleteCart{
		UserSerial: req.UserSerial,
		ID:         req.ID,
	})
}
