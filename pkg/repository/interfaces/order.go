package interfaces

import (
	"GlassGalore/pkg/utils/models"
)

type OrderRepository interface {
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	GetOrders(orderID int) (models.AllItems, error)
	GetAllOrders(userID, page, pageSize int) ([]models.OrderDetails, error)
	CheckOrderStatusByID(id int) (string, error)
	CheckPaymentStatusByID(id int) (string, error)
	CancelOrder(id int) error
	CancelOrderPaid(id int) error
	GetOrderDetailsBrief(page int) ([]models.CombinedOrderDetails, error)
	ChangeOrderStatus(orderID, status string) error
	GetShipmentStatus(orderID string) (string, error)
	// ReturnOrder(id int) error
	GetDetailedOrderThroughId(orderId int) (models.CombinedOrderDetails, error)
	GetOrdersRazor(orderID int) (models.OrderPayOnly, error)
	GetOrderStatus(orderID int) (string, error)
	FindUserID(orderID int) (int, error)
	FindFinalPrice(orderID int) (float64, error)
	ReturnOrder(shipmentStatus string, orderID int) error
	ReduceStockAfterOrder(productName string, quantitry int) error
	GetProductDetailsFromOrder(orderID int) ([]models.OrderProducts, error)
	UpdateQuantityProduct(orderProducts []models.OrderProducts) error
	CheckOrderExist(id int) (bool, error)
}
