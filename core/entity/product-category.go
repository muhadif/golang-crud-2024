package entity

type ProductCategory struct {
	ID     int64  `json:"id"`
	Serial string `json:"serial" gorm:"primaryKey"`
	Name   string `json:"name"`
}

func (p ProductCategory) TableName() string {
	return "product_category"
}

type ProductProductCategory struct {
	ID                    int64  `json:"id"`
	ProductSerial         string `json:"product_serial" gorm:"primaryKey"`
	ProductCategorySerial string `json:"product_category_serial" gorm:"primaryKey"`
}
