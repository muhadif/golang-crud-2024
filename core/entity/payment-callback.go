package entity

type PaymentCallbackVATransferRequest struct {
	TransactionID string `json:"transactionId"`
	Status        PaymentCallbackStatus
}

type PaymentCallbackStatus string

const (
	PaymentCallbackStatusSuccess = "SUCCESS"
	PaymentCallbackStatusFailed  = "FAILED"
)
