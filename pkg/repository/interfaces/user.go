package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type UserRepository interface {
	UserSignUp(models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	GetUserDetails(id int) (models.UserDetailsResponse, error)
	GetAddresses(id int) ([]domain.Address, error)
	CheckIfFirstAddress(id int) bool
	AddAddress(id int, address models.AddAddress, result bool) error
	EditName(id int, name string) error
	EditEmail(id int, Email string) error
	EditPhone(id int, Phone string) error
}
