package domain

import "time"

type Cart struct {
	ID        uint        `json:"id" gorm:"primary_key;column:id;type:BIGINT AUTO_INCREMENT NOT NULL;"`
	UserID    uint        `json:"user_id" gorm:"column:user_id; NOT NULL;"`
	User      User        `json:"user" gorm:"foreignkey:UserID"`       // FK dari users
	CartItems []CartItems `json:"cart_items" gorm:"foreignkey:CartID"` // FK ke cart_items
	CreatedAt time.Time   `json:"created_at" gorm:"type:DATETIME;column:created_at; autoCreateTime"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"type:DATETIME;column:updated_at; autoUpdateTime"`
}

func (Cart) TableName() string {
	return "carts"
}
