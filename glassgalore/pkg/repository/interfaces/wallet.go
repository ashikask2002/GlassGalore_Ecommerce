package interfaces

import "GlassGalore/pkg/utils/models"

type WalletRepository interface {
	GetWallet(id int) (models.WalletAmount, error)
	FindWalletIDFromUserID(userID int) (int, error)
	CreateNewWallet(userID int) (int, error)
	CreditToUserWallet(finalPrice float64, walletID int) error
}
