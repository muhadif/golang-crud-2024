package repository

import (
	"context"
	"golang-crud-2024/core/entity"
)

type PaymentHistoryRepository interface {
	CreatePayment(ctx context.Context, req *entity.PaymentHistory) error
	GetPaymentHistory(ctx context.Context, userSerial string) ([]*entity.PaymentHistory, error)
	GetPaymentHistoryByTransactionID(ctx context.Context, trxID string) (*entity.PaymentHistory, error)
	UpdatePaymentStatus(ctx context.Context, paymentSerial string, status entity.PaymentStatus) error
}
