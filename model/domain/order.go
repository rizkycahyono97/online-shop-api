package domain

import "time"

type Order struct {
	ID          uint      `json:"id" gorm:"primary_key;column:id; AUTO_INCREMENT;"`
	UserID      uint      `json:"user_id" gorm:"column:user_id; NOT NULL;"`
	User        User      `json:"user" gorm:"foreignkey:UserID"`
	TotalAmount float64   `json:"total_amount" gorm:"column:total_amount; NOT NULL;"`
	Status      string    `json:"status" gorm:"column:status; NOT NULL"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:DATETIME;column:created_at; autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:DATETIME;column:updated_at; autoUpdateTime"`

	OrderItems []OrderItem `json:"order_items" gorm:"OrderID"` // FK ke table order
	Payment    *Payment    `json:"payment" gorm:"foreignkey:OrderID"`
}

func (Order) TableName() string {
	return "orders"
}
