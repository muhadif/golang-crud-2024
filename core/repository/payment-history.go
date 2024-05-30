package repository

import (
	"context"
	"golang-crud-2024/core/entity"
)

type PaymentHistoryRepository interface {
	CreatePayment(ctx context.Context, req *entity.PaymentHistory) error
	GetPaymentHistory(ctx context.Context, userSerial string) (*entity.PaymentHistory, error)
}
