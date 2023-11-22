package interfaces

type CartRepository interface {
	GetCartId(user_id int) (int, error)
	CreateNewCart(user_id int) (int, error)
	CheckIfItemsIsAlreadyAdded(cart_id, inventory_id int) (bool, error)
	AddLineItems(cart_id, inventory_id int) error
}
