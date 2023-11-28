package usecase

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"strconv"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase) services.OrderUseCase {
	return &orderUseCase{
		orderRepository: repo,
		userUseCase:     userUseCase,
	}
}

func (i *orderUseCase) OrderItemsFromCart(userID int, addressID int, paymentID int) error {
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

	// Place an order with the order repository
	orderID, err := i.orderRepository.OrderItems(userID, addressID, paymentID, total)
	if err != nil {
		return err
	}

	// Add individual items from the cart to the order
	if err := i.orderRepository.AddOrderProducts(orderID, cart.Data); err != nil {
		return err
	}

	// Remove the ordered items from the user's cart
	for _, v := range cart.Data {
		if err := i.userUseCase.RemoveFromCart(cart.ID, v.ID); err != nil {
			return err
		}
	}

	return nil
}

func (i *orderUseCase) GetOrders(orderid int) (domain.OrderResponse, error) {

	orders, err := i.orderRepository.GetOrders(orderid)
	if err != nil {
		return domain.OrderResponse{}, err
	}
	return orders, nil
}

func (i *orderUseCase) GerAllOrders(userId, page, pageSize int) ([]models.OrderDetails, error) {
	allOrder, err := i.orderRepository.GetAllOrders(userId, page, pageSize)
	if err != nil {
		return []models.OrderDetails{}, err
	}
	return allOrder, nil
}

func (i *orderUseCase) CancelOrder(orderID int) error {
	orderStatus, err := i.orderRepository.CheckOrderStatusByID(orderID)
	if err != nil {
		return err
	}
	if orderStatus != "PENDING" {
		return errors.New("order cannot be  cancelled , kindly return the product if accidently booked")
	}

	err = i.orderRepository.CancelOrder(orderID)
	if err != nil {
		return err
	}

	return nil
}

func (i *orderUseCase) GetAdminOrders(page int) ([]models.CombinedOrderDetails, error) {
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
