package interfaces

import "GlassGalore/pkg/utils/models"

type InvnetoryUseCase interface {
	AddInventory(inventory models.AddInventories) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	UpdateInventory(ProductID int, Stock int) (models.InventoryResponse, error)
	EditInventoryDetails(int, models.EditInventoryDetails) error
}
