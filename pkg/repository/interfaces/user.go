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
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
	GetCartID(id int) (int, error)
	GetProductsInCart(cart_id int) ([]int, error)
	FindProductNames(inventory_id int) (string, error)
	FindCartQuantity(cart_id, inventory_id int) (int, error)
	FindPrice(inventory_id int) (float64, error)
	FindCategory(inventory_id int) (int, error)
	FindStock(id int) (int, error)
	RemoveFromCart(cart, inventory int) error
	UpdateQuantity(id, inv_id, qty int) error
}
