package usecase

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase services.UserUseCase) *orderUseCase {
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
