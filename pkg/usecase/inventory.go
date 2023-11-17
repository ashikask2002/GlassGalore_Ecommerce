package usecase

import (
	helper_interface "GlassGalore/pkg/helper/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"

	interfaces "GlassGalore/pkg/repository/interfaces"
)

type invnetoryUseCase struct {
	repository interfaces.InventoryRepository
	helper     helper_interface.Helper
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, h helper_interface.Helper) *invnetoryUseCase {
	return &invnetoryUseCase{
		repository: repo,
		helper:     h,
	}
}

func (i *invnetoryUseCase) AddInventory(inventory models.AddInventories) (models.InventoryResponse, error) {
	InventoryResponse, err := i.repository.AddInventory(inventory)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return InventoryResponse, nil
}

func (i *invnetoryUseCase) DeleteInventory(inventoryID string) error {
	err := i.repository.DeleteInventory(inventoryID)
	if err != nil {
		return err
	}
	return nil
}

func (i *invnetoryUseCase) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {
	result, err := i.repository.CheckInventory(pid)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	if !result {
		return models.InventoryResponse{}, errors.New("there is not inventory as you mentioned")
	}
	newcat, err := i.repository.UpdateInventory(pid, stock)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	return newcat, err
}

func (i *invnetoryUseCase) EditInventoryDetails(id int, model models.EditInventoryDetails) error {
	//send the url and save it in to the database

	err := i.repository.EditInventoryDetails(id, model)
	if err != nil {
		return err
	}
	return nil
}
