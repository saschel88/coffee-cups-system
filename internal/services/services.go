package services

import (
	"gorm.io/gorm"
)

// Services holds all service dependencies
type Services struct {
	User    *UserService
	Coffee  *CoffeeService
	Box     *BoxService
	Payment *PaymentService
}

// NewServices creates a new Services instance with all dependencies
func NewServices(db *gorm.DB, logger interface{}) *Services {
	return &Services{
		User:    NewUserService(db),
		Coffee:  NewCoffeeService(db),
		Box:     NewBoxService(db),
		Payment: NewPaymentService(db),
	}
}
