package interfaces

import (
	"GlassGalore/pkg/utils/models"

	"github.com/jung-kurt/gofpdf"
)

type OrderUseCase interface {
	OrderItemsFromCart(userID int, addressID int, paymentID int, couponID int) error
	GetOrders(orderid int) (models.AllItems, error)
	GerAllOrders(UserId, page, pageSize int) ([]models.OrderDetails, error)
	CancelOrder(orderID int) error
	GetAdminOrders(page int) ([]models.CombinedOrderDetails, error)
	OrdersStatus(orderID string) error
	// ReturnOrder(id int) error
	ReturnOrder(orderId int) error
	PrintInvoice(orderIdInt int) (*gofpdf.Fpdf, error)
}
