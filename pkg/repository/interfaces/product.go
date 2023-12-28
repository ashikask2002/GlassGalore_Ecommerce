package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type ProductRepository interface {
	AddProduct(inventory models.AddProducts) (domain.Products, error)
	DeleteProduct(id string) error
	CheckProduct(pid int) (bool, error)
	UpdateProduct(pid int, stock int) (models.ProductResponse, error)
	EditProductDetails(id int, model models.EditProductDetails) (models.EditProductDetails, error)
	ListProducts(page int) ([]models.Products, error)
	CheckStock(inventory_id int) (int, error)
	SearchProducts(offset, limit int, search string) ([]models.ProductUserResponse, error)
	FilterProducts(CategroyID int) ([]models.ProductUserResponse, error)
	FilterProductsByPrice(Price, pricetwo int) ([]models.ProductUserResponse, error)
	GetCatOffer(id int) (float64, error)
	CheckIfProductAlreadyExists(productname string) (bool, error)
	GetIdExist(id int) (bool, error)
	Rating(id, productid int, rating float64) error
	FindRating(id int) (float64, error)
	UpdateProductImage(productID int,url string) error
}
