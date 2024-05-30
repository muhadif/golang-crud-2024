package module

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
)

type PaymentCallbackService interface {
	CallbackPaymentVATransfer(ctx context.Context, req *entity.PaymentCallbackVATransferRequest) error
}

type paymentCallbackService struct {
	paymentHistoryRepo repository.PaymentHistoryRepository
}

func NewPaymentCallbackService(paymentHistoryRepo repository.PaymentHistoryRepository) PaymentCallbackService {
	return paymentCallbackService{paymentHistoryRepo: paymentHistoryRepo}
}

func (p paymentCallbackService) CallbackPaymentVATransfer(ctx context.Context, req *entity.PaymentCallbackVATransferRequest) error {
	paymentHistory, err := p.paymentHistoryRepo.GetPaymentHistoryByTransactionID(ctx, req.TransactionID)
	if err != nil {
		return err
	}

	err = p.paymentHistoryRepo.UpdatePaymentStatus(ctx, paymentHistory.Serial, parsePaymentCallbackStatusToPaymentStatus(req.Status))
	if err != nil {
		return err
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
