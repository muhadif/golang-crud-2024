package module

import (
	"context"
	"golang-crud-2024/core/entity"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/fault"
	string2 "golang-crud-2024/pkg/string"
	"time"
)

type PaymentHistoryService interface {
	CreatePayment(ctx context.Context, req *entity.CreatePaymentRequest) (*entity.PaymentHistory, error)
	GetPaymentHistory(ctx context.Context, userSerial string) ([]*entity.PaymentHistory, error)
	GetPaymentBySerial(ctx context.Context, req *entity.GetPaymentBySerialRequest) (*entity.PaymentHistory, error)
	CancelPaymentBySerial(ctx context.Context, req *entity.CancelPaymentBySerialRequest) error
}

type paymentHistory struct {
	paymentHistoryRepo repository.PaymentHistoryRepository
	checkoutService    CheckoutService
	productRepository  repository.ProductRepository
}

func NewPaymentHistoryService(paymentHistoryRepo repository.PaymentHistoryRepository, checkoutService CheckoutService, productRepository repository.ProductRepository) PaymentHistoryService {
	return &paymentHistory{paymentHistoryRepo: paymentHistoryRepo, checkoutService: checkoutService, productRepository: productRepository}
}

func (p paymentHistory) CreatePayment(ctx context.Context, req *entity.CreatePaymentRequest) (*entity.PaymentHistory, error) {
	currentCheckoutItem, err := p.checkoutService.GetCurrentCheckout(ctx, req.UserSerial)
	if err != nil {
		return nil, err
	}

	if len(currentCheckoutItem.CartItems) == 0 {
		return nil, fault.ErrorDictionary(fault.HTTPBadRequestError, coreErr.ErrCheckoutCartIsEmpty)
	}

	paymentSerial := string2.GenerateSerial(entity.PaymentSerialPrefix, 5)

	var paymentItem []*entity.PaymentHistoryItem
	for _, item := range currentCheckoutItem.CartItems {
		item.Product.Stock -= item.Quantity

		paymentItem = append(paymentItem, &entity.PaymentHistoryItem{
			PaymentHistorySerial: paymentSerial,
			ProductSerial:        item.ProductSerial,
			Price:                item.Product.Price,
			Quantity:             item.Quantity,
			Product:              item.Product,
		})
	}

	currentTime := time.Now()

	transactionID := string2.GenerateTransactionID()

	paymentHistory := &entity.PaymentHistory{
		Serial:        paymentSerial,
		OpenTime:      &currentTime,
		UserSerial:    req.UserSerial,
		TotalPrice:    currentCheckoutItem.Total,
		PaymentItems:  paymentItem,
		PaymentMethod: req.PaymentMethod,
		TransactionID: transactionID,
		Status:        entity.PaymentStatusWaiting,
	}

	err = p.paymentHistoryRepo.CreatePayment(ctx, paymentHistory)
	if err != nil {
		return nil, err
	}

	err = p.checkoutService.DeleteCheckout(ctx, req.UserSerial)
	if err != nil {
		return nil, err
	}
	return paymentHistory, nil
}

func (p paymentHistory) GetPaymentHistory(ctx context.Context, userSerial string) ([]*entity.PaymentHistory, error) {
	return p.paymentHistoryRepo.GetPaymentHistory(ctx, userSerial)
}

func (p paymentHistory) GetPaymentBySerial(ctx context.Context, req *entity.GetPaymentBySerialRequest) (*entity.PaymentHistory, error) {
	return p.paymentHistoryRepo.GetPaymentBySerial(ctx, req)
}

func (p paymentHistory) CancelPaymentBySerial(ctx context.Context, req *entity.CancelPaymentBySerialRequest) error {
	paymentHistory, err := p.paymentHistoryRepo.GetPaymentBySerial(ctx, &entity.GetPaymentBySerialRequest{
		PaymentSerial: req.PaymentSerial,
		UserSerial:    req.UserSerial,
	})
	if err != nil {
		return err
	}

	if paymentHistory.Status != entity.PaymentStatusWaiting {
		return fault.ErrorDictionary(fault.HTTPForbiddenRequestError, coreErr.ErrPaymentNotAcceptable)
	}

	err = p.paymentHistoryRepo.UpdatePaymentStatus(ctx, paymentHistory.Serial, entity.PaymentStatusCancelled)
	if err != nil {
		return err
	}

	for _, paymentItem := range paymentHistory.PaymentItems {
		if err := p.productRepository.RollbackStock(ctx, &entity.RollbackStockRequest{
			ProductSerial: paymentItem.ProductSerial,
			RollbackStock: paymentItem.Quantity,
		}); err != nil {
			return err
		}
	}

	return nil
}
