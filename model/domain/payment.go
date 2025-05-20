package domain

import "time"

type Payment struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null;uniqueIndex"`
	UserID    uint      `json:"user_id"`
	Method    string    `json:"payment_method" gorm:"column:payment_method"`
	Amount    float64   `json:"amount" gorm:"column:amount"`
	Status    string    `json:"status" gorm:"column:orders_status"`
	IsPaid    bool      `json:"is_paid" gorm:"column:is_paid;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (Payment) TableName() string {
	return "payments"
}
