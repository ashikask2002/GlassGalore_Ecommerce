package interfaces

import "GlassGalore/pkg/utils/models"

type UserRepository interface {
	UserSignUp(models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
}
