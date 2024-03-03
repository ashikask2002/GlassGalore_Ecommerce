package usecase

import (
	"GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
)

type walletUsecases struct {
	walletRepository interfaces.WalletRepository
}

func NewWalletUseCase(repository interfaces.WalletRepository) services.WalletUseCase {
	return &walletUsecases{
		walletRepository: repository,
	}
}

func (i *walletUsecases) GetWallet(id int) (models.WalletAmount, error) {
	return i.walletRepository.GetWallet(id)
}
