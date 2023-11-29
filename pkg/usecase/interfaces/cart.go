package interfaces

import "GlassGalore/pkg/utils/models"

type CartUseCase interface {
	AddToCart(user_id, product_id int) error
	CheckOut(id int) (models.CheckOut, error)
}
