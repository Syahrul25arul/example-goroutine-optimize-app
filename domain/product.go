package domain

type Product struct {
	ProductId   int64  `gorm:"PrimaryKey" json:"product_id" required:"false"`
	ProductName string `db:"product_name" json:"product_name" required:"true"`
	CategoryId  int64  `db:"category_id" json:"category_id" required:"true"`
}

func (Product) TableName() string {
	return "products"
}
