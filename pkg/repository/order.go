package repository

import (
	"GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"fmt"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{

		DB: db,
	}
}

func (i *orderRepository) OrderItems(userid, addresid, paymentid int, total float64) (int, error) {
	fmt.Println("gggggg", addresid, total)
	var id int
	query := `
	 INSERT INTO orders (created_at,user_id,address_id,payment_method_id,final_price)
	 VALUES (Now(),?, ?, ?, ?)
	 RETURNING id`

	i.DB.Raw(query, userid, addresid, paymentid, total).Scan(&id)

	return id, nil

}

func (i *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {
	fmt.Println("zxzxzxzx", order_id)
	query := `
    INSERT INTO order_items (order_id,inventory_id,quantity,total_price)
    VALUES (?, ?, ?, ?)
    `
	for _, v := range cart {
		var inv int
		if err := i.DB.Raw("select id from inventories where product_name = $1", v.ProductName).Scan(&inv).Error; err != nil {
			return err
		}
		if err := i.DB.Exec(query, order_id, inv, v.Qantity, v.Total).Error; err != nil {
			return err
		}
	}
	return nil
}
