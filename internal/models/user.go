package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	TelegramID int64          `json:"telegram_id" gorm:"uniqueIndex;not null"`
	Username   string         `json:"username"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	IsActive   bool           `json:"is_active" gorm:"default:true"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	CoffeeLogs []CoffeeLog `json:"coffee_logs,omitempty" gorm:"foreignKey:UserID"`
	Payments   []Payment   `json:"payments,omitempty" gorm:"foreignKey:UserID"`
}

// TableName returns the table name for User
func (User) TableName() string {
	return "users"
}
