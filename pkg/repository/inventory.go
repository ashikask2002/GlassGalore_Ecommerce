package repository

import (
	"GlassGalore/pkg/utils/models"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) *inventoryRepository {
	return &inventoryRepository{
		DB: DB,
	}
}

func (i *inventoryRepository) AddInventory(inventory models.AddInventories) (models.InventoryResponse, error) {

	query := `INSERT INTO inventories (category_id, product_name, size, stock, price) VALUES (?, ?, ?, ?, ?);`

	err := i.DB.Exec(query, inventory.CategoryID, inventory.ProductName, inventory.Size, inventory.Stock, inventory.Price).Error
	if err != nil {
		return models.InventoryResponse{}, err
	}

	var inventoryResponse models.InventoryResponse

	return inventoryResponse, nil
}

func (i *inventoryRepository) DeleteInventory(inventoryID string) error {
	id, err := strconv.Atoi(inventoryID)
	if err != nil {
		return errors.New("Converting to integer not happened")
	}

	result := i.DB.Exec("DELETE FROM inventories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no results with that id exist")
	}
	return nil
}

func (i *inventoryRepository) CheckInventory(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM inventories WHERE id = ?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}
	if k == 0 {
		return false, err
	}
	return true, err
}

func (i *inventoryRepository) UpdateInventory(pid int, stock int) (models.InventoryResponse, error) {
	//check database connection
	if i.DB == nil {
		return models.InventoryResponse{}, errors.New("database connection is nill")
	}

	//update the inventories
	if err := i.DB.Exec("UPDATE inventories SET stock = stock + $1 WHERE id = $2", stock, pid).Error; err != nil {
		return models.InventoryResponse{}, err
	}

	//retrieve the update
	var newDetails models.InventoryResponse
	var newStock int
	if err := i.DB.Raw("SELECT stock FROM inventories WHERE id = ?", pid).Scan(&newStock).Error; err != nil {
		return models.InventoryResponse{}, err
	}
	newDetails.ProductID = pid
	newDetails.Stock = stock

	return newDetails, nil
}

func (i *inventoryRepository) EditInventoryDetails(id int, model models.EditInventoryDetails) error {
	err := i.DB.Exec("UPDATE inventories SET product_name = $1, category_id = $2, price = $3, size = $4 WHERE id =$5", model.Name, model.CategoryID, model.Price, model.Size, id).Error
	if err != nil {
		return err
	}
	return nil
}
