package usecase

import (
	"GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type orderUseCase struct {
	orderRepository  interfaces.OrderRepository
	userUseCase      services.UserUseCase
	walletRepository interfaces.WalletRepository
	couponRepository interfaces.CouponRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase, walletRepository interfaces.WalletRepository, couponRepository interfaces.CouponRepository) services.OrderUseCase {
	return &orderUseCase{
		orderRepository:  repo,
		userUseCase:      userUseCase,
		walletRepository: walletRepository,
		couponRepository: couponRepository,
	}
}

func (i *orderUseCase) OrderItemsFromCart(userID int, addressID int, paymentID int, couponID int) error {
	if userID <= 0 {
		return errors.New("userid must be positive")
	}
	if addressID <= 0 {
		return errors.New("address id must be positive")
	}
	if couponID <= 0 {
		return errors.New("coupon id must be positive")
	}
	if paymentID <= 0 {
		return errors.New("payment id must be positive")
	}
	// Retrieve the user's cart
	cart, err := i.userUseCase.GetCart(userID)
	if err != nil {
		return err
	}

	// Calculate the total price of the items in the cart
	var total float64
	for _, item := range cart.Data {
		if item.Quantity > 0 && item.Price > 0 {
			total += float64(item.Quantity) * float64(item.Price)
		}
	}

	couponvalid, err := i.couponRepository.CheckCouponValid(couponID)
	if err != nil {
		return err
	}
	if !couponvalid {
		return errors.New("this coupon is invalid")
	}

	coupon, err := i.couponRepository.FindCouponPrice(couponID)
	if err != nil {
		return err
	}

	totaldiscount := float64(coupon)

	total = total - totaldiscount
	if total <= 0 {
		return errors.New("there is nothing in your cart")
	}

	// Place an order with the order repository
	orderID, err := i.orderRepository.OrderItems(userID, addressID, paymentID, total)
	if err != nil {
		return err
	}

	// Add individual items from the cart to the order
	if err := i.orderRepository.AddOrderProducts(orderID, cart.Data); err != nil {
		return err
	}
	//decrease the stock of product after order
	for _, v := range cart.Data {
		if err := i.orderRepository.ReduceStockAfterOrder(v.ProductName, v.Quantity); err != nil {
			return err
		}
	}

	// Remove the ordered items from the user's cart
	for _, v := range cart.Data {
		if err := i.userUseCase.RemoveFromCart(cart.ID, v.ID); err != nil {
			return err
		}
	}

	return nil
}

func (i *orderUseCase) GetOrders(orderid int) (models.AllItems, error) {

	orders, err := i.orderRepository.GetOrders(orderid)
	if err != nil {
		return models.AllItems{}, err
	}
	return orders, nil
}

func (i *orderUseCase) GerAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error) {

	fmt.Println("page is ", page)
	allOrder, err := i.orderRepository.GetAllOrders(userId, page, pageSize)
	if err != nil {
		return []models.OrderDetails{}, err
	}
	return allOrder, nil
}

func (i *orderUseCase) CancelOrder(orderID int) error {
	// Check the order status
	if orderID <= 0 {
		return errors.New("orderid must be positive")
	}
	orderStatus, err := i.orderRepository.CheckOrderStatusByID(orderID)
	if err != nil {
		return err
	}
	if orderStatus == "RETURNED" {
		return errors.New("the order already returned")
	}

	if orderStatus == "CANCELED" {
		return errors.New("the order already canelled")
	}

	orderProductDetails, err := i.orderRepository.GetProductDetailsFromOrder(orderID)
	if err != nil {
		return err
	}

	// Check if the order can be canceled
	if orderStatus == "PENDING" || orderStatus == "SHIPPED" {

		paymentStatus, err := i.orderRepository.CheckPaymentStatusByID(orderID)
		if err != nil {
			return err
		}
		if paymentStatus != "PAID" {
			err = i.orderRepository.CancelOrder(orderID)
			if err != nil {
				return err
			}
			err = i.orderRepository.UpdateQuantityProduct(orderProductDetails)
			if err != nil {
				return err
			}
		}
		if paymentStatus == "PAID" {
			UserID, err := i.orderRepository.FindUserID(orderID)
			if err != nil {
				return err
			}

			FinalPrice, err := i.orderRepository.FindFinalPrice(orderID)
			if err != nil {
				return err
			}

			//find if the user having the wallet

			walletID, err := i.walletRepository.FindWalletIDFromUserID(UserID)
			if err != nil {
				return err
			}

			//if not any wallet create a new wallet

			if walletID == 0 {
				walletID, err = i.walletRepository.CreateNewWallet(UserID)
				if err != nil {
					return err
				}
			}

			//add the price to the wallet
			if err := i.walletRepository.CreditToUserWallet(FinalPrice, walletID); err != nil {
				return err
			}
			// Attempt to cancel the order
			err = i.orderRepository.CancelOrderPaid(orderID)
			if err != nil {
				return err
			}

			err = i.orderRepository.UpdateQuantityProduct(orderProductDetails)
			if err != nil {
				return err
			}

		}

	}
	return nil

}

func (i *orderUseCase) GetAdminOrders(page int) ([]models.CombinedOrderDetails, error) {
	if page <= 0 {
		return []models.CombinedOrderDetails{}, errors.New("page must be positive")
	}
	orderDetails, err := i.orderRepository.GetOrderDetailsBrief(page)
	if err != nil {
		return []models.CombinedOrderDetails{}, err
	}
	return orderDetails, nil
}

func (i *orderUseCase) OrdersStatus(orderID string) error {
	orderId, sErr := strconv.Atoi(orderID)
	if sErr != nil {
		return sErr
	}
	if orderId <= 0 {
		return errors.New("order id must be positive")
	}
	exist, err := i.orderRepository.CheckOrderExist(orderId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("this order is not exist")
	}
	status, err := i.orderRepository.CheckOrderStatusByID(orderId)
	if err != nil {
		return err
	}
	switch status {
	case "CANCELED", "RETURNED", "DELIVERED":
		return errors.New("cannot approve this order becouse it's in processed or cancelled state")
	case "PENDING":
		//for the admin approval change the PENDING to SHIPPED
		err := i.orderRepository.ChangeOrderStatus(orderID, "SHIPPED")
		if err != nil {
			return err
		}
	case "SHIPPED":
		shipmentStatus, err := i.orderRepository.GetShipmentStatus(orderID)
		if err != nil {
			return err
		}
		if shipmentStatus == "CANCELED" {
			return errors.New("cannot approve this orders becouse its cancelled")
		}

		//for admin approval, change SHIPPED to  DELIVERED
		err = i.orderRepository.ChangeOrderStatus(orderID, "DELIVERED")
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *orderUseCase) ReturnOrder(orderId int) error {

	if orderId <= 0 {
		return errors.New("orderID must be positive")
	}
	shipmentStatus, err := i.orderRepository.GetOrderStatus(orderId)
	if err != nil {
		return err
	}
	if shipmentStatus == "RETURNED" {
		return errors.New("the order already returned")
	}

	if shipmentStatus != "DELIVERED" {
		return errors.New("user try to return the order that not even delivered")
	}

	if shipmentStatus == "DELIVERED" {

		UserID, err := i.orderRepository.FindUserID(orderId)
		if err != nil {
			return err
		}

		FinalPrice, err := i.orderRepository.FindFinalPrice(orderId)
		if err != nil {
			return err
		}

		//find if the user having the wallet

		walletID, err := i.walletRepository.FindWalletIDFromUserID(UserID)
		if err != nil {
			return err
		}

		//if not any wallet create a new wallet

		if walletID == 0 {
			walletID, err = i.walletRepository.CreateNewWallet(UserID)
			if err != nil {
				return err
			}
		}

		//add the price to the wallet
		if err := i.walletRepository.CreditToUserWallet(FinalPrice, walletID); err != nil {
			return err
		}

		if err := i.orderRepository.ReturnOrder("RETURNED", orderId); err != nil {
			return err
		}
		return nil
	}

	return errors.New("cannot return order")
}

func (i *orderUseCase) PrintInvoice(orderId int) (*gofpdf.Fpdf, error) {
	if orderId <= 0 {
		return nil, errors.New("order id must be positive")
	}
	order, err := i.orderRepository.GetDetailedOrderThroughId(orderId)
	if err != nil {
		return nil, err
	}

	items, err := i.orderRepository.GetItemsByOrderID(orderId)
	if err != nil {
		return nil, err
	}

	if order.OrderStatus != "DELIVERED" {
		return nil, errors.New("wait for the invoice until the product is delivered")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 30)
	pdf.SetTextColor(31, 73, 125)
	pdf.Cell(0, 20, "Invoice")
	pdf.Ln(20)

	pdf.SetFont("Arial", "I", 14)
	pdf.SetTextColor(51, 51, 51)
	pdf.Cell(0, 10, "Customer Details")
	pdf.Ln(10)
	customerDetails := []string{
		"Name: " + order.Name,
		"House Name: " + order.HouseName,
		"Street: " + order.Street,
		"State: " + order.State,
		"City: " + order.City,
	}
	for _, detail := range customerDetails {
		pdf.Cell(0, 10, detail)
		pdf.Ln(10)
	}
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(40, 10, "Item", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Price", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Quantity", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Final Price", "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.SetFillColor(255, 255, 255)
	for _, item := range items {
		pdf.CellFormat(40, 10, item.ProductName, "1", 0, "L", true, 0, "")
		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Price, 'f', 2, 64), "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, strconv.Itoa(item.Quantity), "1", 0, "C", true, 0, "")
		pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(item.Price*float64(item.Quantity), 'f', 2, 64), "1", 0, "C", true, 0, "")
		pdf.Ln(10)
	}
	pdf.Ln(10)

	var totalPrice float64
	for _, item := range items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.CellFormat(120, 10, "Total Price:", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(totalPrice, 'f', 2, 64), "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	offerApplied := totalPrice - order.FinalPrice

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.CellFormat(120, 10, "Offer Applied:", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(offerApplied, 'f', 2, 64), "1", 0, "C", true, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 16)
	pdf.SetFillColor(217, 217, 217)
	pdf.CellFormat(120, 10, "Final Amount:", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, "$"+strconv.FormatFloat(order.FinalPrice, 'f', 2, 64), "1", 0, "C", true, 0, "")
	pdf.Ln(10)
	pdf.SetFont("Arial", "I", 12)
	pdf.Cell(0, 10, "Generated by Crocsclub India Pvt Ltd. - "+time.Now().Format("2006-01-02 15:04:05"))
	pdf.Ln(10)

	return pdf, nil
}
