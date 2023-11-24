package interfaces

type OrderUseCase interface {
	OrderItemsFromCart(userID int, addressID int, paymentID int) error
}
