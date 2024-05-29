package entity

type Product struct {
	ID          int64   `json:"id"`
	Serial      string  `json:"serial"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}
