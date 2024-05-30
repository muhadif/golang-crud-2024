package module

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	string2 "golang-crud-2024/pkg/string"
	"time"
)

type PaymentHistoryService interface {
	CreatePayment(ctx context.Context, req *entity.CreatePaymentRequest) error
	GetPaymentHistory(ctx context.Context, userSerial string) (*entity.PaymentHistory, error)
}

type paymentHistory struct {
	paymentHistoryRepo repository.PaymentHistoryRepository
	checkoutService    CheckoutService
}

func NewPaymentHistoryService(paymentHistoryRepo repository.PaymentHistoryRepository, checkoutService CheckoutService) PaymentHistoryService {
	return &paymentHistory{paymentHistoryRepo: paymentHistoryRepo, checkoutService: checkoutService}
}

func (p paymentHistory) CreatePayment(ctx context.Context, req *entity.CreatePaymentRequest) error {
	currentCheckoutItem, err := p.checkoutService.GetCurrentCheckout(ctx, req.UserSerial)
	if err != nil {
		return err
	}

	paymentSerial := string2.GenerateSerial(entity.PaymentSerialPrefix, 5)

	var paymentItem []*entity.PaymentHistoryItem
	for _, item := range currentCheckoutItem.CartItems {
		paymentItem = append(paymentItem, &entity.PaymentHistoryItem{
			PaymentHistorySerial: paymentSerial,
			ProductSerial:        item.ProductSerial,
			Price:                item.Product.Price,
			Quantity:             item.Quantity,
			Product:              item.Product,
		})
	}

	currentTime := time.Now()

	paymentHistory := &entity.PaymentHistory{
		Serial:        paymentSerial,
		OpenTime:      &currentTime,
		UserSerial:    req.UserSerial,
		TotalPrice:    currentCheckoutItem.Total,
		PaymentItems:  paymentItem,
		PaymentMethod: req.PaymentMethod,
		Status:        entity.PaymentStatusWaiting,
	}

	err = p.paymentHistoryRepo.CreatePayment(ctx, paymentHistory)
	if err != nil {
		return err
	}

	err = p.checkoutService.DeleteCheckout(ctx, req.UserSerial)
	if err != nil {
		return err
	}
	return nil
}

func (p paymentHistory) GetPaymentHistory(ctx context.Context, userSerial string) (*entity.PaymentHistory, error) {
	return nil, nil
}
