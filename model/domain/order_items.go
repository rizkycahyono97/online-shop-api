package domain

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primary_key;column:id; AUTO_INCREMENT;"`
	OrderID   uint    `json:"order_id" gorm:"column:order_id; NOT NULL;"`
	Order     Order   `json:"order" gorm:"foreignkey:OrderID"`
	ProductID uint    `json:"product_id" gorm:"column:product_id; NOT NULL;"`
	Product   Product `json:"product" gorm:"foreignkey:ProductID"`
	Quantity  uint    `json:"quantity" gorm:"column:quantity; NOT NULL;"`
	Price     float32 `json:"price" gorm:"column:price; NOT NULL;"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
