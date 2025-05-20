package domain

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primary_key;column:id;type:BIGINT AUTO_INCREMENT NOT NULL;"`
	Name        string    `json:"name" gorm:"column:name;type:VARCHAR(255);column:name; NOT NULL;"`
	Description string    `json:"description" gorm:"column:description;type:TEXT; NOT NULL;"`
	Price       float64   `json:"price" gorm:"column:price;type:DECIMAL(10, 2); NOT NULL;"`
	Stock       uint      `json:"stock" gorm:"column:stock;type:INT; NOT NULL;"`
	ImageURL    string    `json:"image_url" gorm:"column:image_url;type:VARCHAR(255);"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:DATETIME;column:created_at; autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:DATETIME;column:updated_at; autoUpdateTime"`

	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}

// nama table
func (Product) TableName() string {
	return "products"
}
