package entity

import "time"

type PaymentItem struct {
	ProductSerial         string
	CheckoutSessionSerial string
	Price                 float64
	Quantity              int32
	Product               *Product
}

type PaymentHistory struct {
	Date          *time.Time
	UserSerial    string
	TotalPrice    float64
	PaymentItems  []*PaymentItem
	PaymentMethod PaymentMethod
	PaymentStatus PaymentStatus
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
