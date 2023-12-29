package repository

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"database/sql"
	"errors"
	"fmt"
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

func (i *productRepository) AddProduct(product models.AddProducts) (domain.Products, error) {
	query := `
		INSERT INTO products (category_id, product_name, discription, size, stock, price)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id, category_id, product_name, discription, size, stock, price;`

	var insertedProduct domain.Products
	err := i.DB.Raw(query,
		product.CategoryID, product.ProductName, product.Discription, product.Size, product.Stock, product.Price).
		Scan(&insertedProduct).Error

	if err != nil {
		return domain.Products{}, err
	}

	fmt.Println("Inserted Product:", insertedProduct)
	return insertedProduct, nil
}

func (i *productRepository) DeleteProduct(productID string) error {
	id, err := strconv.Atoi(productID)
	if err != nil {
		return errors.New("converting to integer not happened")
	}
	if id <= 0 {
		return errors.New("id must be positive")
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
	fmt.Println("yyyyyyyyyyyyyyyyyyy", id, model)
	err := i.DB.Exec("UPDATE products SET product_name = $1, discription = $2, category_id = $3, price = $4, size = $5 WHERE id =$6", model.Name, model.Discription, model.CategoryID, model.Price, model.Size, id).Error
	if err != nil {
		return models.EditProductDetails{}, err
	}
	var products models.EditProductDetails
	err = i.DB.Raw("SELECT product_name as name,discription,category_id,price,size FROM products WHERE id = ?", id).Scan(&products).Error
	if err != nil {
		return models.EditProductDetails{}, err
	}
	fmt.Println("fkfkjkfjdkjfkdjkfjdkjfkjdk", products)
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

func (i *productRepository) FilterProducts(CategoryID int) ([]models.ProductUserResponse, error) {
	var product_list []models.ProductUserResponse

	if err := i.DB.Raw("select * from products where category_id = ? ", CategoryID).Scan(&product_list).Error; err != nil {
		return nil, err
	}
	return product_list, nil
}

func (i *productRepository) FilterProductsByPrice(Price, pricetwo int) ([]models.ProductUserResponse, error) {
	var product_list []models.ProductUserResponse

	if err := i.DB.Raw("SELECT * FROM products WHERE price BETWEEN ? AND ?", Price, pricetwo).Scan(&product_list).Error; err != nil {
		return nil, err
	}
	return product_list, nil
}

func (i *productRepository) SearchProducts(offset, limit int, search string) ([]models.ProductUserResponse, error) {
	var product_list []models.ProductUserResponse

	query := "SELECT id, product_name, price, category_id, size FROM products WHERE product_name LIKE ? LIMIT ? OFFSET ?"
	err := i.DB.Raw(query, search+"%", limit, offset).Scan(&product_list).Error

	if err != nil {
		return nil, errors.New("record not found")
	}
	fmt.Println("vzvzvzvzvzvzvzv", product_list)

	return product_list, nil
}

func (i *productRepository) GetCatOffer(id int) (float64, error) {
	var disc_price float64
	if err := i.DB.Raw("select discount_price from category_offers where category_id = ?", id).Scan(&disc_price).Error; err != nil {
		return 0, err
	}
	return disc_price, nil
}
func (i *productRepository) CheckIfProductAlreadyExists(productname string) (bool, error) {
	var count int64
	err := i.DB.Raw("SELECT COUNT(*) FROM products WHERE product_name = ?", productname).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (i *productRepository) GetIdExist(id int) (bool, error) {
	var count int64
	err := i.DB.Raw("SELECT COUNT(*) FROM products WHERE category_id = ?", id).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (i *productRepository) Rating(id, prductid int, rating float64) error {
	result := i.DB.Exec("INSERT INTO ratings (user_id, product_id, rating) VALUES (?, ?, ?)", id, prductid, rating)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (i *productRepository) FindRating(id int) (float64, error) {
	var rating sql.NullFloat64
	err := i.DB.Raw("SELECT COALESCE(AVG(CASE WHEN rating IS NOT NULL THEN rating ELSE 0 END), 0) FROM ratings WHERE product_id = ?", id).Scan(&rating).Error
	if err != nil {
		return 0, err
	}
	return rating.Float64, nil
}

func (i *productRepository) UpdateProductImage(productID int, url string) error {
	err := i.DB.Exec("insert into images (product_id,url) values($1,$2) returning *", productID, url).Error
	if err != nil {
		return errors.New("error while insert image to database")
	}
	return nil

}

func (i *productRepository) GetImage(productID int) (string, error) {
	var url string
	err := i.DB.Raw("select url from images where product_id = ?", productID).Scan(&url).Error
	if err != nil {
		return "", err
	}
	return url, nil
}
