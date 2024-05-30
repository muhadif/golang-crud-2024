package payment_history

import (
	"context"
	"golang-crud-2024/core/entity"
	coreErr "golang-crud-2024/core/error"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/fault"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewRepository(db *gorm.DB) repository.PaymentHistoryRepository {
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

func (r repo) CreatePayment(ctx context.Context, req *entity.PaymentHistory) (err error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(req).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, item := range req.PaymentItems {
		var product *entity.Product
		err := tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).Where("serial = ?", item.ProductSerial).First(&product).Error
		if err != nil {
			return err
		}

		if product.Stock < item.Quantity {
			err = fault.ErrorDictionary(fault.HTTPBadRequestError, coreErr.ErrProductStock)
			return err
		}

		err = tx.WithContext(ctx).Exec("UPDATE product SET stock = stock - ? WHERE serial = ?", item.Quantity, item.Product.Serial).Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Create(item).Error
		if err != nil {
			return err
		}
	}

	return nil

}

func (r repo) GetPaymentHistory(ctx context.Context, userSerial string) ([]*entity.PaymentHistory, error) {
	var paymentHistories []*entity.PaymentHistory

	err := r.db.WithContext(ctx).Where("user_serial = ?", userSerial).Order("created_at DESC").Find(&paymentHistories).Error

	if err != nil {
		return nil, err
	}

	return paymentHistories, nil
}

func (r repo) GetPaymentHistoryByTransactionID(ctx context.Context, trxID string) (*entity.PaymentHistory, error) {
	var payment *entity.PaymentHistory

	err := r.db.WithContext(ctx).Where("transaction_id", trxID).First(&payment).Error
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r repo) UpdatePaymentStatus(ctx context.Context, paymentSerial string, status entity.PaymentStatus) error {
	err := r.db.WithContext(ctx).Model(&entity.PaymentHistory{}).
		Where("serial = ?", paymentSerial).
		Update("status", status).Error

	if err != nil {
		return err
	}

	return nil
}
