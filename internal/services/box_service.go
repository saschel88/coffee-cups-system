package services

import (
	"fmt"

	"github.com/your-username/coffee-cups-system/internal/models"
	"gorm.io/gorm"
)

// BoxService handles box-related operations
type BoxService struct {
	db *gorm.DB
}

// NewBoxService creates a new BoxService
func NewBoxService(db *gorm.DB) *BoxService {
	return &BoxService{db: db}
}

// CreateBox creates a new coffee box
func (s *BoxService) CreateBox(name string, totalCups int, price float64, createdBy uint) (*models.Box, error) {
	box := models.Box{
		Name:      name,
		TotalCups: totalCups,
		Price:     price,
		CreatedBy: createdBy,
		IsActive:  true,
	}

	if err := s.db.Create(&box).Error; err != nil {
		return nil, fmt.Errorf("failed to create box: %w", err)
	}

	return &box, nil
}

// GetActiveBoxes retrieves all active boxes
func (s *BoxService) GetActiveBoxes() ([]models.Box, error) {
	var boxes []models.Box
	err := s.db.Where("is_active = ?", true).Preload("Creator").Find(&boxes).Error
	return boxes, err
}

// GetBoxByID retrieves a box by ID
func (s *BoxService) GetBoxByID(id uint) (*models.Box, error) {
	var box models.Box
	err := s.db.Preload("Creator").First(&box, id).Error
	if err != nil {
		return nil, err
	}
	return &box, nil
}

// DeactivateBox deactivates a box
func (s *BoxService) DeactivateBox(id uint) error {
	return s.db.Model(&models.Box{}).Where("id = ?", id).Update("is_active", false).Error
}
