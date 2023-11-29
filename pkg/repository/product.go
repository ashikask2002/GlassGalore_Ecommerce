package repository

import (
	"GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productRepository{
		DB: DB,
	}
}

func (i *productRepository) AddProduct(product models.AddProducts) (models.ProductResponse, error) {

	query := `INSERT INTO products (category_id, product_name, size, stock, price) VALUES (?, ?, ?, ?, ?);`

	err := i.DB.Exec(query, product.CategoryID, product.ProductName, product.Size, product.Stock, product.Price).Error
	if err != nil {
		return models.ProductResponse{}, err
	}

	var productResponse models.ProductResponse

	return productResponse, nil
}

func (i *productRepository) DeleteProduct(productID string) error {
	id, err := strconv.Atoi(productID)
	if err != nil {
		return errors.New("converting to integer not happened")
	}

	result := i.DB.Exec("DELETE FROM products WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no results with that id exist")
	}
	return nil
}

func (i *productRepository) CheckProduct(pid int) (bool, error) {
	var k int
	err := i.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", pid).Scan(&k).Error
	if err != nil {
		return false, err
	}
	if k == 0 {
		return false, err
	}
	return true, err
}

func (i *productRepository) UpdateProduct(pid int, stock int) (models.ProductResponse, error) {
	//check database connection
	if i.DB == nil {
		return models.ProductResponse{}, errors.New("database connection is nill")
	}

	//update the products
	if err := i.DB.Exec("UPDATE products SET stock = stock + $1 WHERE id = $2", stock, pid).Error; err != nil {
		return models.ProductResponse{}, err
	}

	//retrieve the update
	var newDetails models.ProductResponse
	var newStock int
	if err := i.DB.Raw("SELECT stock FROM products WHERE id = ?", pid).Scan(&newStock).Error; err != nil {
		return models.ProductResponse{}, err
	}
	newDetails.ProductID = pid
	newDetails.Stock = stock

	return newDetails, nil
}

func (i *productRepository) EditProductDetails(id int, model models.EditProductDetails) (models.EditProductDetails, error) {
	err := i.DB.Exec("UPDATE products SET product_name = $1, category_id = $2, price = $3, size = $4 WHERE id =$5", model.Name, model.CategoryID, model.Price, model.Size, id).Error
	if err != nil {
		return models.EditProductDetails{}, err
	}
	var products models.EditProductDetails
	err = i.DB.Raw("SELECT product_name as name,category_id,price,size FROM products WHERE id = ?", id).Scan(&products).Error
	if err != nil {
		return models.EditProductDetails{}, err
	}
	return products, nil
}

func (i *productRepository) ListProducts(page int) ([]models.Products, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 10

	var productDetails []models.Products

	if err := i.DB.Raw("select id,category_id,product_name,size,stock,price from products limit $1 offset $2", 10, offset).Scan(&productDetails).Error; err != nil {
		return []models.Products{}, err
	}

	return productDetails, nil
}

func (i *productRepository) CheckStock(pid int) (int, error) {
	var k int
	if err := i.DB.Raw("select stock from products where id = ?", pid).Scan(&k).Error; err != nil {
		return 0, err
	}
	return k, nil

}
