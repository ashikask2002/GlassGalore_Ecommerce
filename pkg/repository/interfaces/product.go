package interfaces

import "GlassGalore/pkg/utils/models"

type ProductRepository interface {
	AddProduct(inventory models.AddProducts) (models.ProductResponse, error)
	DeleteProduct(id string) error
	CheckProduct(pid int) (bool, error)
	UpdateProduct(pid int, stock int) (models.ProductResponse, error)
	EditProductDetails(id int, model models.EditProductDetails) (models.EditProductDetails, error)
	ListProducts(page int) ([]models.Products, error)
	CheckStock(inventory_id int) (int, error)
}
