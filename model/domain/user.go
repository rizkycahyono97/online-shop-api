package domain

import "time"

type User struct {
	ID        uint64    `json:"id" gorm:"primaryKey;column:id;type:BIGINT AUTO_INCREMENT NOT NULL;"`
	Name      string    `json:"name" gorm:"type:VARCHAR(255);column:name; NOT NULL;"`
	Email     string    `json:"email" gorm:"type:VARCHAR(255);column:email; NOT NULL;"`
	Password  string    `json:"password" gorm:"type:VARCHAR(255);column:password; NOT NULL;"`
	Role      string    `json:"role" gorm:"type:enum('customer','admin');default:'customer;column:role; NOT NULL;'"`
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME;column:created_at; autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:DATETIME;column:updated_at; autoUpdateTime"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;type:DATETIME;index"`
}

// nama table
func (User) TableName() string {
	return "users"
}
