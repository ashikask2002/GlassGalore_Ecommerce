package usecase

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/helper"
	helper_interface "GlassGalore/pkg/helper/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"mime/multipart"

	repos "GlassGalore/pkg/repository/interfaces"
	use "GlassGalore/pkg/usecase/interfaces"
)

type productUseCase struct {
	repository repos.ProductRepository
	helper     helper_interface.Helper
}

func NewProductUseCase(repo repos.ProductRepository, h helper_interface.Helper) use.ProductUseCase {
	return &productUseCase{
		repository: repo,
		helper:     h,
	}
}

func (i *productUseCase) AddProduct(Products models.AddProducts) (domain.Products, error) {
	if Products.CategoryID <= 0 {
		return domain.Products{}, errors.New("id must be positive")
	}
	if Products.Discription == "" {
		return domain.Products{}, errors.New("discription must be something")
	}
	if Products.Price <= 0 {
		return domain.Products{}, errors.New("price must be positive")
	}
	if Products.ProductName == "" {
		return domain.Products{}, errors.New("product name is empty now")
	}
	if Products.Stock <= 0 {
		return domain.Products{}, errors.New("stock must be positive")
	}

	exists, err := i.repository.CheckIfProductAlreadyExists(Products.ProductName)
	if err != nil {
		return domain.Products{}, err
	}
	if exists {
		return domain.Products{}, errors.New("product is already exist")
	}
	productResponse, err := i.repository.AddProduct(Products)
	if err != nil {
		return domain.Products{}, err
	}
	return productResponse, nil
}

func (i *productUseCase) DeleteProduct(productID string) error {

	err := i.repository.DeleteProduct(productID)
	if err != nil {
		return err
	}
	return nil
}

func (i *productUseCase) UpdateProduct(pid int, stock int) (models.ProductResponse, error) {

	if pid <= 0 || stock <= 0 {
		return models.ProductResponse{}, errors.New("no negative values are allowded")
	}
	result, err := i.repository.CheckProduct(pid)
	if err != nil {
		return models.ProductResponse{}, err
	}
	if !result {
		return models.ProductResponse{}, errors.New("there is not product as you mentioned")
	}
	newcat, err := i.repository.UpdateProduct(pid, stock)
	if err != nil {
		return models.ProductResponse{}, err
	}
	return newcat, err
}

func (i *productUseCase) EditProductDetails(id int, model models.EditProductDetails) (models.EditProductDetails, error) {
	//send the url and save it in to the database
	if id <= 0 {
		return models.EditProductDetails{}, errors.New("id must be positive")
	}
	if model.CategoryID <= 0 {
		return models.EditProductDetails{}, errors.New("category id must be positive")
	}
	if model.Name == "" {
		return models.EditProductDetails{}, errors.New("name is empty now")
	}
	if model.Price <= 0 {
		return models.EditProductDetails{}, errors.New("price must be positive")
	}
	if model.Discription == "" {
		return models.EditProductDetails{}, errors.New("discription must have something")
	}

	products, err := i.repository.EditProductDetails(id, model)
	if err != nil {
		return models.EditProductDetails{}, err
	}
	return products, nil
}

func (i *productUseCase) ListProductForUser(page int) ([]models.Products, error) {
	if page <= 0 {
		return []models.Products{}, errors.New("page must be positive")
	}
	productDetails, err := i.repository.ListProducts(page)
	if err != nil {
		return []models.Products{}, err
	}

	for j := range productDetails {
		discountPrice, err := i.repository.GetCatOffer(productDetails[j].CategoryID)
		if err != nil {
			return []models.Products{}, errors.New("error in getting category offer")
		}

		productDetails[j].DiscountPrice = productDetails[j].Price - discountPrice
	}

	for k := range productDetails {
		productRating, err := i.repository.FindRating(int(productDetails[k].ID))
		if err != nil {
			return []models.Products{}, errors.New("error in getting the rating")
		}
		productDetails[k].Rating = productRating
	}
	var UpdatedProductDetails []models.Products

	for _, p := range productDetails {
		img, err := i.repository.GetImage(int(p.ID))
		if err != nil {
			return nil, err
		}
		p.Image = img
		UpdatedProductDetails = append(UpdatedProductDetails, p)
	}

	return UpdatedProductDetails, nil
}

func (i *productUseCase) FilterProducts(CategoryID int) ([]models.ProductUserResponse, error) {
	if CategoryID <= 0 {
		return []models.ProductUserResponse{}, errors.New("you provided the negtive id")
	}

	exist, err := i.repository.GetIdExist(CategoryID)
	if err != nil {
		return []models.ProductUserResponse{}, err
	}
	if !exist {
		return []models.ProductUserResponse{}, errors.New("not exist")
	}

	product_list, err := i.repository.FilterProducts(CategoryID)
	if err != nil {
		return []models.ProductUserResponse{}, err
	}
	for j := range product_list {
		discountPrice, err := i.repository.GetCatOffer(product_list[j].CategoryID)
		if err != nil {
			return []models.ProductUserResponse{}, errors.New("error in getting category offer")
		}

		product_list[j].OfferPrice = product_list[j].Price - discountPrice
	}

	return product_list, nil
}

func (i *productUseCase) FilterProductsByPrice(Price, pricetwo int) ([]models.ProductUserResponse, error) {
	if Price <= 0 {
		return []models.ProductUserResponse{}, errors.New("you provided the negtive price")
	}
	if pricetwo <= 0 {
		return []models.ProductUserResponse{}, errors.New("you proivded the negtive price")
	}
	product_list, err := i.repository.FilterProductsByPrice(Price, pricetwo)
	if err != nil {
		return []models.ProductUserResponse{}, err
	}

	for j := range product_list {
		discountPrice, err := i.repository.GetCatOffer(product_list[j].CategoryID)
		if err != nil {
			return []models.ProductUserResponse{}, errors.New("error in getting category offer")
		}

		product_list[j].OfferPrice = product_list[j].Price - discountPrice
	}
	return product_list, nil
}

func (i *productUseCase) SearchProducts(search models.Search) ([]models.ProductUserResponse, error) {

	offset := (search.Page - 1) * search.Limit

	product_list, err := i.repository.SearchProducts(offset, search.Limit, search.Search)
	if err != nil {
		return nil, err
	}
	for j := range product_list {
		discountPrice, err := i.repository.GetCatOffer(product_list[j].CategoryID)
		if err != nil {
			return []models.ProductUserResponse{}, errors.New("error in getting category offer")
		}

		product_list[j].OfferPrice = product_list[j].Price - discountPrice
	}

	return product_list, nil

}

func (i *productUseCase) Rating(id, productid int, rating float64) error {
	err := i.repository.Rating(id, productid, rating)
	if err != nil {
		return err
	}
	return nil
}

func (i *productUseCase) UpdateProductImage(id int, file *multipart.FileHeader) error {
	url, err := helper.AddImageToS3(file)
	if err != nil {
		return err
	}
	err = i.repository.UpdateProductImage(id, url)
	if err != nil {
		return err
	}
	return nil
}
