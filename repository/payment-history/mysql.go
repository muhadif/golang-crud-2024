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
		if tx != nil {
			if err := tx.Commit(); err != nil {
				tx.Rollback()
			}
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

func (r repo) GetPaymentHistory(ctx context.Context, userSerial string) (*entity.PaymentHistory, error) {
	//TODO implement me
	panic("implement me")
}
