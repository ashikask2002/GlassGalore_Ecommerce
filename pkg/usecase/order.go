package usecase

import (
	"GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
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

	cart, err := i.userUseCase.GetCart(userID)
	if err != nil {
		return err
	}

	var total float64
	for _, item := range cart.Data {
		if item.Qantity > 0 && item.Price > 0 {
			total += float64(item.Qantity) * float64(item.Price)
		}
	}

	orderID, err := i.orderRepository.OrderItems(userID, addressID, paymentID, total)
	if err != nil {
		return err
	}

	if err := i.orderRepository.AddOrderProducts(orderID, cart.Data); err != nil {
		return err
	}

	for _, v := range cart.Data {
		if err := i.userUseCase.RemoveFromCart(cart.ID, v.ID); err != nil {
			return err
		}
	}
	return nil
}
