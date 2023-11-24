package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	GetUserByID(id string) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	CheckIfPaymentMethodAlreadyExists(payment string) (bool, error)
	NewPaymentMethod(payment string) error
	ListPaymentMethods() ([]domain.PaymentMethod, error)
	DeletePaymentMethod(id int) error
	GetPaymentMethod() ([]models.PaymentMethodResponse, error)
}
