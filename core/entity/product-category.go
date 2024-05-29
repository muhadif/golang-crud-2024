package entity

type ProductCategory struct {
	ID     int64  `json:"id"`
	Serial string `json:"serial"`
	Name   string `json:"name"`
}

type ProductCategoryMapping struct {
	ID                    int64  `json:"id"`
	ProductSerial         string `json:"product_serial"`
	ProductCategorySerial string `json:"product_category_serial"`
}
