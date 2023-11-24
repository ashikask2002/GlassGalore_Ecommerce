package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type OrderRepository interface {
	OrderItems(userid, addressid, paymentid int, total float64) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	GetOrders(orderid int) (domain.OrderResponse, error)
	GetAllOrders(userID, page, pageSize int) ([]models.OrderDetails, error)
	CheckOrderStatusByID(id int) (string, error)
	CancelOrder(id int) error
}
