package services

import (
	"fmt"

	"github.com/your-username/coffee-cups-system/internal/models"
	"gorm.io/gorm"
)

// UserService handles user-related operations
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// CreateOrUpdateUser creates a new user or updates an existing one
func (s *UserService) CreateOrUpdateUser(telegramID int64, username, firstName, lastName string) (*models.User, error) {
	var user models.User
	err := s.db.Where("telegram_id = ?", telegramID).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		// Create new user
		user = models.User{
			TelegramID: telegramID,
			Username:   username,
			FirstName:  firstName,
			LastName:   lastName,
			IsActive:   true,
		}
		if err := s.db.Create(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	} else {
		// Update existing user
		user.Username = username
		user.FirstName = firstName
		user.LastName = lastName
		if err := s.db.Save(&user).Error; err != nil {
			return nil, fmt.Errorf("failed to update user: %w", err)
		}
	}

	return &user, nil
}

// GetUserByTelegramID retrieves a user by their Telegram ID
func (s *UserService) GetUserByTelegramID(telegramID int64) (*models.User, error) {
	var user models.User
	err := s.db.Where("telegram_id = ?", telegramID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllActiveUsers retrieves all active users
func (s *UserService) GetAllActiveUsers() ([]models.User, error) {
	var users []models.User
	err := s.db.Where("is_active = ?", true).Find(&users).Error
	return users, err
}
