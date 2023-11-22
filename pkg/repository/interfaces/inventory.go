package interfaces

import "GlassGalore/pkg/utils/models"

type InventoryRepository interface {
	AddInventory(inventory models.AddInventories) (models.InventoryResponse, error)
	DeleteInventory(id string) error
	CheckInventory(pid int) (bool, error)
	UpdateInventory(pid int, stock int) (models.InventoryResponse, error)
	EditInventoryDetails(id int, model models.EditInventoryDetails) error
	ListProducts(page int) ([]models.Inventories, error)
	CheckStock(inventory_id int) (int, error)
}
