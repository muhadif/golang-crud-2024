package entity

type Product struct {
	ID                int64             `json:"id" `
	Serial            string            `json:"serial" gorm:"primaryKey"`
	Name              string            `json:"name"`
	Price             float64           `json:"price"`
	Stock             int32             `json:"stock"`
	Description       string            `json:"description"`
	ProductCategories []ProductCategory `json:"productCategories,omitempty" gorm:"many2many:product_product_category;foreignKey:serial;joinForeignKey:product_serial;References:serial;joinReferences:product_category_serial"`
}

func (p Product) TableName() string {
	return "product"
}

type GetProductRequest struct {
	ProductCategorySerial string `form:"productCategorySerial" `
}
