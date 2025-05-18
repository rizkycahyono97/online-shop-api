package domain

import "time"

type Payment struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint `gorm:"not null;uniqueIndex"`
	Method    string
	Amount    float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Payment) TableName() string {
	return "payments"
}
