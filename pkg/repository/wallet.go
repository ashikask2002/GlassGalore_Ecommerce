package repository

import (
	"GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(DB *gorm.DB) interfaces.WalletRepository {
	return &WalletRepository{
		DB: DB,
	}
}

func (i *WalletRepository) GetWallet(userId int) (models.WalletAmount, error) {
	var WalletAmount models.WalletAmount

	if err := i.DB.Raw("select amount from wallets where user_id = ?", userId).Scan(&WalletAmount).Error; err != nil {
		return models.WalletAmount{}, err
	}
	return WalletAmount, nil

}

func (i *WalletRepository) FindWalletIDFromUserID(userID int) (int, error) {
	var count int
	err := i.DB.Raw("select count(*) from wallets where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	var walletID int
	if count > 0 {
		err := i.DB.Raw("select id from wallets where user_id = ?", userID).Scan(&walletID).Error
		if err != nil {
			return 0, err
		}
	}
	return walletID, nil
}

func (i *WalletRepository) CreateNewWallet(userID int) (int, error) {
	var walletID int
	err := i.DB.Exec("insert into wallets (user_id,amount) values ($1,$2)", userID, 0).Error
	if err != nil {
		return 0, err
	}
	if err := i.DB.Raw("select id from wallets where user_id = ?", userID).Scan(&walletID).Error; err != nil {
		return 0, err
	}
	return walletID, nil
}

func (i *WalletRepository) CreditToUserWallet(finalPrice float64, walletID int) error {

	if err := i.DB.Exec("update wallets set amount = ? where id = ?", finalPrice, walletID).Error; err != nil {
		return err
	}
	return nil
}
