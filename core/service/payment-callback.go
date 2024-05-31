package module

import (
	"context"
	"golang-crud-2024/core/entity"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/fault"
)

type PaymentCallbackService interface {
	CallbackPaymentVATransfer(ctx context.Context, req *entity.PaymentCallbackVATransferRequest) error
}

type paymentCallbackService struct {
	paymentHistoryRepo repository.PaymentHistoryRepository
	productRepository  repository.ProductRepository
}

func NewPaymentCallbackService(paymentHistoryRepo repository.PaymentHistoryRepository, productRepository repository.ProductRepository) PaymentCallbackService {
	return paymentCallbackService{paymentHistoryRepo: paymentHistoryRepo, productRepository: productRepository}
}

func (p paymentCallbackService) CallbackPaymentVATransfer(ctx context.Context, req *entity.PaymentCallbackVATransferRequest) error {
	paymentHistory, err := p.paymentHistoryRepo.GetPaymentHistoryByTransactionID(ctx, req.TransactionID)
	if err != nil {
		return err
	}

	if paymentHistory.Status != entity.PaymentStatusWaiting {
		return fault.ErrorDictionary(fault.HTTPForbiddenRequestError, coreErr.ErrPaymentNotAcceptable)
	}

	err = p.paymentHistoryRepo.UpdatePaymentStatus(ctx, paymentHistory.Serial, parsePaymentCallbackStatusToPaymentStatus(req.Status))
	if err != nil {
		return err
	}

	// TODO implement with transaction
	if req.Status == entity.PaymentCallbackStatusSuccess {
		return nil
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

func parsePaymentCallbackStatusToPaymentStatus(status entity.PaymentCallbackStatus) entity.PaymentStatus {
	switch status {
	case entity.PaymentCallbackStatusSuccess:
		return entity.PaymentStatusPaid
	case entity.PaymentCallbackStatusFailed:
		return entity.PaymentCallbackStatusFailed
	default:
		return entity.PaymentStatusWaiting
	}
}
