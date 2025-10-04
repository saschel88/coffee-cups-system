package models

import (
	"time"

	"gorm.io/gorm"
)

// Box represents a coffee box/capsule package
type Box struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	TotalCups int            `json:"total_cups" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedBy uint           `json:"created_by" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Relationships
	Creator    User        `json:"creator,omitempty" gorm:"foreignKey:CreatedBy"`
	CoffeeLogs []CoffeeLog `json:"coffee_logs,omitempty" gorm:"foreignKey:BoxID"`
	Payments   []Payment   `json:"payments,omitempty" gorm:"foreignKey:BoxID"`
}

// TableName returns the table name for Box
func (Box) TableName() string {
	return "boxes"
}

// GetUsedCups returns the number of cups used from this box
func (b *Box) GetUsedCups(db *gorm.DB) (int, error) {
	var count int64
	err := db.Model(&CoffeeLog{}).Where("box_id = ?", b.ID).Count(&count).Error
	return int(count), err
}

// GetRemainingCups returns the number of remaining cups in this box
func (b *Box) GetRemainingCups(db *gorm.DB) (int, error) {
	used, err := b.GetUsedCups(db)
	if err != nil {
		return 0, err
	}
	return b.TotalCups - used, nil
}
