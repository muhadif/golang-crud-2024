package entity

type Cart struct {
	ID            int64    `json:"id"`
	UserSerial    string   `json:"userSerial"`
	ProductSerial string   `json:"productSerial"`
	Product       *Product `json:"product" gorm:"foreignKey:ProductSerial"`
	Quantity      int32    `json:"quantity"`
}

func (c Cart) TableName() string {
	return "cart"
}

type CreateCart struct {
	UserSerial    string `json:"userSerial"`
	ProductSerial string `json:"productSerial"`
	Quantity      int32  `json:"quantity"`
}

type UpdateCart struct {
	ID            int64  `json:"id"`
	UserSerial    string `json:"userSerial"`
	ProductSerial string `json:"productSerial"`
	Quantity      int32  `json:"quantity"`
}

type DeleteCart struct {
	ID         int64  `json:"id"`
	UserSerial string `json:"userSerial"`
}

type GetCartByID struct {
	ID         string `json:"id"`
	UserSerial string `json:"userSerial"`
}
