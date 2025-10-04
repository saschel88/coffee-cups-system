package models

import (
	"time"

	"gorm.io/gorm"
)

// Payment represents a payment made by a user for coffee consumption
type Payment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	BoxID     uint           `json:"box_id" gorm:"not null;index"`
	Amount    float64        `json:"amount" gorm:"not null"`
	IsPaid    bool           `json:"is_paid" gorm:"default:false"`
	PaidAt    *time.Time     `json:"paid_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Box  Box  `json:"box,omitempty" gorm:"foreignKey:BoxID"`
}

// TableName returns the table name for Payment
func (Payment) TableName() string {
	return "payments"
}
