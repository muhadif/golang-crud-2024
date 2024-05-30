package entity

type GetCheckoutSessionResponse struct {
	CartItems []*Cart
	Total     float64
}

type CreateCheckoutSession struct {
	CartItems  []*CartCheckoutItem `json:"cartItems"`
	UserSerial string              `json:"userSerial"`
}

type CartCheckoutItem struct {
	CartID     int64  `json:"cartId"`
	UserSerial string `json:"userSerial"`
}