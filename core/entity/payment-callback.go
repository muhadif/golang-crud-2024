package entity

type PaymentCallbackVATransferRequest struct {
	TransactionID string                `json:"transactionId"`
	Status        PaymentCallbackStatus `json:"status" binding:"required,enum"`
}

type PaymentCallbackStatus string

const (
	PaymentCallbackStatusSuccess = "SUCCESS"
	PaymentCallbackStatusFailed  = "FAILED"
)

func (s PaymentCallbackStatus) IsValid() bool {
	switch s {
	case PaymentCallbackStatusSuccess, PaymentCallbackStatusFailed:
		return true
	}

	return false
}
