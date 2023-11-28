package repository

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
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
		if err := i.DB.Exec(query, order_id, inv, v.Quantity, v.Total).Error; err != nil {
			return err
		}
	}
	return nil
}

func (i *orderRepository) GetOrders(orderId int) (domain.OrderResponse, error) {
	if orderId <= 0 {
		return domain.OrderResponse{}, errors.New("order id should be positive number")
	}
	var order domain.OrderResponse

	query := `SELECT * FROM orders WHERE id = $1`

	if err := i.DB.Raw(query, orderId).First(&order).Error; err != nil {
		return domain.OrderResponse{}, err
	}
	return order, nil
}

func (i *orderRepository) GetAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	var order []models.OrderDetails

	err := i.DB.Raw("SELECT id as order_id, address_id, payment_method_id, final_price as price, order_status, payment_status FROM  orders WHERE user_id=? OFFSET ? LIMIT ?", userId, offset, pageSize).Scan(&order).Error
	if err != nil {
		return nil, err
	}
	fmt.Println("Retrieved orders:", order)
	return order, nil
}

func (i *orderRepository) CheckOrderStatusByID(id int) (string, error) {
	var status string
	err := i.DB.Raw("select order_status from orders where id = ?", id).Scan(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

func (i *orderRepository) CancelOrder(id int) error {
	if err := i.DB.Exec("update orders set order_status='CANCELED' where id=$1", id).Error; err != nil {
		return err
	}
	return nil
}

func (i *orderRepository) GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 2

	var orderDetails []models.CombinedOrderDetails

	err := i.DB.Raw(`
	SELECT orders.id AS order_id, orders.final_price, orders.order_status, orders.payment_status, 
	users.name, users.email, users.phone, addresses.house_name, addresses.state, 
	addresses.pin, addresses.street, addresses.city 
	FROM orders 
	INNER JOIN users ON orders.user_id = users.id 
	INNER JOIN addresses ON users.id = addresses.user_id 
	LIMIT ? OFFSET ?
`, 2, offset).Scan(&orderDetails).Error

	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}
	return orderDetails, nil

}

func (i *orderRepository) ChangeOrderStatus(orderID, status string) error {
	err := i.DB.Exec("UPDATE orders SET order_status = ? WHERE id = ?", status, orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *orderRepository) GetShipmentStatus(orderID string) (string, error) {

	var shipmentStatus string
	err := i.DB.Raw("UPDATE orders SET order_status = 'DELIVERED', payment_status = 'PAID' WHERE id = ?", orderID).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}
	return shipmentStatus, nil
}
