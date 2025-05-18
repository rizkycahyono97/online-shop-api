package domain

import "time"

type Order struct {
	ID          uint
	UserID      uint
	User        User
	TotalAmount float64
	Status      string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	OrderItems []OrderItem

	Payment Payment `gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "orders"
}
