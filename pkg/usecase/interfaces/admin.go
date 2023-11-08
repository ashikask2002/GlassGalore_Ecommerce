package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers(page int) ([]models.UserDetailsAtAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
}
