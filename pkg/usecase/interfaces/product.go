package interfaces

import "GlassGalore/pkg/utils/models"

type InvnetoryUseCase interface {
	AddProduct(inventory models.AddProducts) (models.ProductResponse, error)
	DeleteProduct(id string) error
	UpdateProduct(ProductID int, Stock int) (models.ProductResponse, error)
	EditProductDetails(int, models.EditProductDetails) (models.EditProductDetails, error)
	ListProductForUser(page int) ([]models.Products, error)
}
