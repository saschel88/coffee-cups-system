package services

import (
	"fmt"
	"time"

	"github.com/your-username/coffee-cups-system/internal/models"
	"gorm.io/gorm"
)

// CoffeeService handles coffee-related operations
type CoffeeService struct {
	db *gorm.DB
}

// NewCoffeeService creates a new CoffeeService
func NewCoffeeService(db *gorm.DB) *CoffeeService {
	return &CoffeeService{db: db}
}

// GetDB returns the database connection
func (s *CoffeeService) GetDB() *gorm.DB {
	return s.db
}

// LogCoffee logs a coffee consumption
func (s *CoffeeService) LogCoffee(userID, boxID uint) (*models.CoffeeLog, error) {
	// Check if the box exists and is active
	var box models.Box
	if err := s.db.Where("id = ? AND is_active = ?", boxID, true).First(&box).Error; err != nil {
		return nil, fmt.Errorf("box not found or inactive: %w", err)
	}

	// Check if there are remaining cups
	remaining, err := box.GetRemainingCups(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to get remaining cups: %w", err)
	}
	if remaining <= 0 {
		return nil, fmt.Errorf("no remaining cups in this box")
	}

	// Create coffee log
	coffeeLog := models.CoffeeLog{
		UserID:   userID,
		BoxID:    boxID,
		LoggedAt: time.Now(),
	}

	if err := s.db.Create(&coffeeLog).Error; err != nil {
		return nil, fmt.Errorf("failed to log coffee: %w", err)
	}

	return &coffeeLog, nil
}

// GetUserCoffeeLogs retrieves coffee logs for a user
func (s *CoffeeService) GetUserCoffeeLogs(userID uint, limit int) ([]models.CoffeeLog, error) {
	var logs []models.CoffeeLog
	query := s.db.Where("user_id = ?", userID).Order("logged_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	err := query.Preload("Box").Find(&logs).Error
	return logs, err
}

// GetBoxStats retrieves statistics for a box
func (s *CoffeeService) GetBoxStats(boxID uint) (*BoxStats, error) {
	var box models.Box
	if err := s.db.Preload("CoffeeLogs").First(&box, boxID).Error; err != nil {
		return nil, err
	}

	used, err := box.GetUsedCups(s.db)
	if err != nil {
		return nil, err
	}

	remaining := box.TotalCups - used
	costPerCup := box.Price / float64(box.TotalCups)

	return &BoxStats{
		Box:           box,
		UsedCups:      used,
		RemainingCups: remaining,
		CostPerCup:    costPerCup,
	}, nil
}

// BoxStats represents statistics for a box
type BoxStats struct {
	Box           models.Box `json:"box"`
	UsedCups      int        `json:"used_cups"`
	RemainingCups int        `json:"remaining_cups"`
	CostPerCup    float64    `json:"cost_per_cup"`
}
