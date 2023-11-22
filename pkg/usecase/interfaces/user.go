package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
	GetUserDetails(id int) (models.UserDetailsResponse, error)
	GetAddresses(id int) ([]domain.Address, error)
	AddAddress(id int, address models.AddAddress) error
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error
	ChangePassword(id int, old string, password string, repassword string) error
	GetCart(id int) (models.GetCartResponse, error)
	RemoveFromCart(cart, inventory int) error
	UpdateQuantity(id, inv_id, qty int) error
}
