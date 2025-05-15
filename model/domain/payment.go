package domain

import "time"

type Payment struct {
	ID        uint      `json:"id" gorm:"primary_key;column:id; AUTO_INCREMENT;"`
	OrderID   uint      `json:"order_id" gorm:"column:order_id; NOT NULL;"`
	Order     Order     `json:"order" gorm:"foreignkey:OrderID"`
	Method    string    `json:"method" gorm:"column:method; NOT NULL;"`
	Amount    float64   `json:"amount" gorm:"column:amount; NOT NULL;"`
	Status    string    `json:"status" gorm:"column:status; NOT NULL;"`
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME;column:created_at; autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:DATETIME;column:updated_at; autoUpdateTime"`
}
