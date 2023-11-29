package usecase

import (
	helper_interface "GlassGalore/pkg/helper/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"

	repos "GlassGalore/pkg/repository/interfaces"
	use "GlassGalore/pkg/usecase/interfaces"
)

type productUseCase struct {
	repository repos.ProductRepository
	helper     helper_interface.Helper
}

func NewProductUseCase(repo repos.ProductRepository, h helper_interface.Helper) use.InvnetoryUseCase {
	return &productUseCase{
		repository: repo,
		helper:     h,
	}
}

func (i *productUseCase) AddProduct(Products models.AddProducts) (models.ProductResponse, error) {
	productResponse, err := i.repository.AddProduct(Products)
	if err != nil {
		return models.ProductResponse{}, err
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

	products, err := i.repository.EditProductDetails(id, model)
	if err != nil {
		return models.EditProductDetails{}, err
	}
	return products, nil
}

func (i *productUseCase) ListProductForUser(page int) ([]models.Products, error) {
	productDetails, err := i.repository.ListProducts(page)
	if err != nil {
		return []models.Products{}, err
	}

	return productDetails, nil
}
