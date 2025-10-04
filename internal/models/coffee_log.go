package models

import (
	"time"

	"gorm.io/gorm"
)

// CoffeeLog represents a coffee consumption log entry
type CoffeeLog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	BoxID     uint           `json:"box_id" gorm:"not null;index"`
	LoggedAt  time.Time      `json:"logged_at" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Box  Box  `json:"box,omitempty" gorm:"foreignKey:BoxID"`
}

// TableName returns the table name for CoffeeLog
func (CoffeeLog) TableName() string {
	return "coffee_logs"
}
