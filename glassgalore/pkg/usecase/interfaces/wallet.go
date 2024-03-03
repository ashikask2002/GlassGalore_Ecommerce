package interfaces

import "GlassGalore/pkg/utils/models"

type WalletUseCase interface {
	GetWallet(id int) (models.WalletAmount, error)
}
