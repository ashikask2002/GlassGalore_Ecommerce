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
    INSERT INTO order_items (order_id,product_id,quantity,total_price)
    VALUES (?, ?, ?, ?)
    `
	for _, v := range cart {
		var inv int
		if err := i.DB.Raw("select id from products where product_name = $1", v.ProductName).Scan(&inv).Error; err != nil {
			return err
		}
		if err := i.DB.Exec(query, order_id, inv, v.Quantity, v.Total).Error; err != nil {
			return err
		}
	}
	return nil
}

func (i *orderRepository) GetOrders(orderID int) (models.OrderPay, error) {
	if orderID <= 0 {
		return models.OrderPay{}, errors.New("order ID should be a positive number")
	}

	fmt.Println("order ID:", orderID)

	var order models.OrderPay

	query := `SELECT * FROM orders WHERE id = $1`
	fmt.Println("abcd", models.OrderPay{})
	// fmt.Println("abcd", domain.Order.AddressID)
	// fmt.Println("abcd", domain.Order.AddressID)
	// fmt.Println("abcd", domain.Order.AddressID)

	if err := i.DB.Raw(query, orderID).Scan(&order).Error; err != nil {
		return models.OrderPay{}, err

	}
	fmt.Println("abcd", domain.Order{})

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

func (i *orderRepository) GetDetailedOrderThroughId(orderId int) (models.CombinedOrderDetails, error) {
	fmt.Println("dddddddreoiii", orderId)
	var body models.CombinedOrderDetails

	query := `
	SELECT 
        o.id AS order_id,
        o.final_price AS final_price,
        o.order_status AS order_status,
        o.payment_status AS payment_status,
        u.name AS name,
        u.email AS email,
        u.phone AS phone,
        a.house_name AS house_name,
        a.state AS state,
        a.pin AS pin,
        a.street AS street,
        a.city AS city
	FROM orders o
	JOIN users u ON o.user_id = u.id
	JOIN addresses a ON o.address_id = a.id 
	WHERE o.id = ?
	`
	if err := i.DB.Raw(query, orderId).Scan(&body).Error; err != nil {
		err = errors.New("error in getting detailed order through id in repository: " + err.Error())
		return models.CombinedOrderDetails{}, err
	}
	fmt.Println("order", body.OrderID)
	return body, nil
}

func (i *orderRepository) GetOrdersRazor(orderID int) (models.OrderPayOnly, error) {
	if orderID <= 0 {
		return models.OrderPayOnly{}, errors.New("order ID should be a positive number")
	}

	fmt.Println("order ID:", orderID)

	var order models.OrderPayOnly

	query := `SELECT id as order_id,final_price FROM orders WHERE id = $1`
	fmt.Println("abcd", models.OrderPay{})

	if err := i.DB.Raw(query, orderID).Scan(&order).Error; err != nil {
		return models.OrderPayOnly{}, err

	}
	// fmt.Println("abcd", domain.Order{})

	return order, nil
}

func (i *orderRepository) GetOrderStatus(orderId int) (string, error) {
	var shipmentStatus string
	err := i.DB.Raw("SELECT order_status FROM orders WHERE id = ?", orderId).Scan(&shipmentStatus).Error
	if err != nil {
		return "", err
	}
	return shipmentStatus, nil
}

// func (i *orderRepository) FindUserID(orderID int) (int, error) {
// 	var UserID int
// 	err := i.DB.Raw("select user_id from orders where id = ?", orderID).Scan(UserID).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	return UserID, nil
// }

//	func (i *orderRepository) FindFinalPrice(orderID int) (int, error) {
//		var finalprice int
//		err := i.DB.Raw("select final_price from orders where id = ?", orderID).Scan(finalprice).Error
//		if err != nil {
//			return 0, err
//		}
//		return finalprice, nil
//	}
func (i *orderRepository) ReturnOrder(shipmentStatus string, orderID int) error {
	err := i.DB.Exec("update orders set order_status = $1, payment_status ='ReturnToWallet' where id = $2", shipmentStatus, orderID).Error
	if err != nil {
		return err
	}
	return nil
}
