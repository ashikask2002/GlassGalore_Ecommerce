package interfaces

import (
	"GlassGalore/pkg/utils/models"
)

type OrderRepository interface {
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	GetOrders(orderID int) (models.OrderPay, error)
	GetAllOrders(userID, page, pageSize int) ([]models.OrderDetails, error)
	CheckOrderStatusByID(id int) (string, error)
	CancelOrder(id int) error
	GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error)
	ChangeOrderStatus(orderID, status string) error
	GetShipmentStatus(orderID string) (string, error)
	// ReturnOrder(id int) error
	GetDetailedOrderThroughId(orderId int) (models.CombinedOrderDetails, error)
	GetOrdersRazor(orderID int) (models.OrderPayOnly, error) 
	GetOrderStatus(orderID int) (string,error)
	// FindUserID(orderID int) (int, error)
	// FindFinalPrice(orderID int) (int, error)
	ReturnOrder(shipmentStatus string,orderID int) error
}
