package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
	"mime/multipart"
)

type ProductUseCase interface {
	AddProduct(inventory models.AddProducts) (domain.Products, error)
	DeleteProduct(id string) error
	UpdateProduct(ProductID int, Stock int) (models.ProductResponse, error)
	EditProductDetails(int, models.EditProductDetails) (models.EditProductDetails, error)
	ListProductForUser(page int) ([]models.Products, error)
	SearchProducts(search models.Search) ([]models.ProductUserResponse, error)
	FilterProducts(categoryID int) ([]models.ProductUserResponse, error)
	FilterProductsByPrice(Price,pricetwo int) ([]models.ProductUserResponse, error)
	Rating(id,productid int, rating float64) error
	UpdateProductImage(id int,file *multipart.FileHeader) error
}
