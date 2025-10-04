package services

import (
	"fmt"
	"time"

	"github.com/your-username/coffee-cups-system/internal/models"
	"gorm.io/gorm"
)

// PaymentService handles payment-related operations
type PaymentService struct {
	db *gorm.DB
}

// NewPaymentService creates a new PaymentService
func NewPaymentService(db *gorm.DB) *PaymentService {
	return &PaymentService{db: db}
}

// CalculateUserDebt calculates the debt for a user for a specific box
func (s *PaymentService) CalculateUserDebt(userID, boxID uint) (float64, error) {
	// Get the box
	var box models.Box
	if err := s.db.First(&box, boxID).Error; err != nil {
		return 0, fmt.Errorf("box not found: %w", err)
	}

	// Count user's coffee logs for this box
	var count int64
	if err := s.db.Model(&models.CoffeeLog{}).Where("user_id = ? AND box_id = ?", userID, boxID).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count coffee logs: %w", err)
	}

	// Calculate debt
	costPerCup := box.Price / float64(box.TotalCups)
	debt := float64(count) * costPerCup

	return debt, nil
}

// CreatePayment creates a payment record
func (s *PaymentService) CreatePayment(userID, boxID uint, amount float64) (*models.Payment, error) {
	payment := models.Payment{
		UserID: userID,
		BoxID:  boxID,
		Amount: amount,
		IsPaid: false,
	}

	if err := s.db.Create(&payment).Error; err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return &payment, nil
}

// MarkPaymentAsPaid marks a payment as paid
func (s *PaymentService) MarkPaymentAsPaid(paymentID uint) error {
	now := time.Now()
	return s.db.Model(&models.Payment{}).Where("id = ?", paymentID).Updates(map[string]interface{}{
		"is_paid": true,
		"paid_at": &now,
	}).Error
}

// GetUserPayments retrieves payments for a user
func (s *PaymentService) GetUserPayments(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := s.db.Where("user_id = ?", userID).Preload("Box").Find(&payments).Error
	return payments, err
}

// GetBoxPayments retrieves all payments for a box
func (s *PaymentService) GetBoxPayments(boxID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := s.db.Where("box_id = ?", boxID).Preload("User").Find(&payments).Error
	return payments, err
}
