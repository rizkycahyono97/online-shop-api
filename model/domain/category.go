package domain

import "time"

type Category struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	CategoryName string    `gorm:"type:varchar(255);unique;not null" json:"category_name"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:DATETIME;column:created_at; autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:DATETIME;column:updated_at; autoUpdateTime"`
}

func (Category) TableName() string {
	return "categories"
}
