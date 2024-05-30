package entity

import "time"

const (
	PaymentSerialPrefix = "PAY-"
)

type PaymentHistoryItem struct {
	PaymentHistorySerial string
	ProductSerial        string
	Price                float64
	Quantity             int32
	Product              *Product
}

func (r PaymentHistoryItem) TableName() string {
	return "payment_history_item"
}

type PaymentHistory struct {
	Serial        string
	OpenTime      *time.Time
	ExpiredTime   *time.Time
	UserSerial    string
	TotalPrice    float64
	PaymentItems  []*PaymentHistoryItem `gorm:"-"`
	PaymentMethod PaymentMethod
	Status        PaymentStatus
}

func (r PaymentHistory) TableName() string {
	return "payment_history"
}

type PaymentMethod string

const (
	VATransfer = "VA_TRANSFER"
)

type PaymentStatus string

const (
	PaymentStatusWaiting   = "WAITING"
	PaymentStatusPaid      = "PAID"
	PaymentStatusExpired   = "EXPIRED"
	PaymentStatusCancelled = "CANCELLED"
)

type CreatePaymentRequest struct {
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	UserSerial    string
}
