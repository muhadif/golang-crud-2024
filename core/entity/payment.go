package entity

import "time"

const (
	PaymentSerialPrefix = "PAY-"
)

type PaymentHistoryItem struct {
	PaymentHistorySerial string   `json:"paymentHistorySerial"`
	ProductSerial        string   `json:"productSerial"`
	Price                float64  `json:"price"`
	Quantity             int32    `json:"quantity"`
	Product              *Product `json:"product" gorm:"foreignKey:ProductSerial"`
}

func (r PaymentHistoryItem) TableName() string {
	return "payment_history_item"
}

type PaymentHistory struct {
	Serial        string                `json:"serial"`
	OpenTime      *time.Time            `json:"openTime"`
	ExpiredTime   *time.Time            `json:"expiredTime"`
	UserSerial    string                `json:"userSerial"`
	TotalPrice    float64               `json:"totalPrice"`
	PaymentItems  []*PaymentHistoryItem `gorm:"-"`
	PaymentMethod PaymentMethod         `json:"paymentMethod"`
	Status        PaymentStatus         `json:"status"`
	TransactionID string                `json:"transactionId"`
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
