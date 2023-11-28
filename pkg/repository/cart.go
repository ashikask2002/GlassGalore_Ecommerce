package repository

import (
	"GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: db,
	}
}

func (i *cartRepository) GetCartId(user_id int) (int, error) {
	var id int
	if err := i.DB.Raw("SELECT id FROM carts WHERE user_id = ?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (i *cartRepository) CreateNewCart(user_id int) (int, error) {
	var id int
	err := i.DB.Exec(`INSERT INTO carts (user_id) VALUES ($1)`, user_id).Error
	if err != nil {
		return 0, err
	}
	if err := i.DB.Raw("select id from carts where user_id = ?", user_id).Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}
func (i *cartRepository) CheckIfItemsIsAlreadyAdded(cart_id, inventory_id int) (bool, error) {
	var count int
	if err := i.DB.Raw("SELECT COUNT(*) FROM line_items WHERE cart_id = $1 AND inventory_id = $2", cart_id, inventory_id).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (i *cartRepository) AddLineItems(cart_id, inventory_id int) error {
	err := i.DB.Exec(`INSERT INTO line_items (cart_id,inventory_id) VALUES ($1,$2)`, cart_id, inventory_id).Error
	if err != nil {
		return err
	}

	return nil
}

func (i *cartRepository) GetAddresses(id int) ([]models.Address, error) {
	var addresses []models.Address
	if err := i.DB.Raw("select * from addresses where user_id = ?", id).Scan(&addresses).Error; err != nil {
		return []models.Address{}, err
	}
	return addresses, nil
}
