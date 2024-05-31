package entity

type GetCheckoutSessionResponse struct {
	CartItems []*Cart `json:"cartItems"`
	Total     float64 `json:"total"`
}

type CreateCheckoutSession struct {
	CartItems  []*CartCheckoutItem `json:"cartItems"`
	UserSerial string              `json:"userSerial"`
}

type CartCheckoutItem struct {
	CartID     int64  `json:"cartId"`
	UserSerial string `json:"userSerial"`
}
