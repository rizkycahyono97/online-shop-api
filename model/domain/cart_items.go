package domain

import "time"

type CartItems struct {
	ID        uint      `json:"id" gorm:"primary_key;column:id;type:BIGINT AUTO_INCREMENT NOT NULL;"`
	CartID    uint      `json:"cart_id" gorm:"NOT NULL"`
	Cart      Cart      `json:"cart" gorm:"foreignkey:CartID"` // FK ke cart
	ProductID uint      `json:"product_id" gorm:"NOT NULL"`
	Product   Product   `json:"product" gorm:"foreignkey:ProductID"` // FK dari product
	Quantity  uint      `json:"quantity" gorm:"column:quantity; NOT NULL"`
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME;column:created_at; autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:DATETIME;column:updated_at; autoUpdateTime"`
}

func (CartItems) TableName() string {
	return "cart_items"
}
