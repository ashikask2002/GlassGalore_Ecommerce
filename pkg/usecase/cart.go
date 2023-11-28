package usecase

import (
	interfaces "GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
)

type cartUseCase struct {
	repo                interfaces.CartRepository
	inventoryRepository interfaces.InventoryRepository
	userUsecase         services.UserUseCase
	adrepo              interfaces.AdminRepository
}

func NewCartUseCase(repo interfaces.CartRepository, inventoryRepo interfaces.InventoryRepository, userUseCase services.UserUseCase, adrepo interfaces.AdminRepository) services.CartUseCase {
	return &cartUseCase{
		repo:                repo,
		inventoryRepository: inventoryRepo,
		userUsecase:         userUseCase,
		adrepo:              adrepo,
	}
}

func (i *cartUseCase) AddToCart(userID, inventoryID int) error {

	//check the desired product has quantity available
	stock, err := i.inventoryRepository.CheckStock(inventoryID)
	if err != nil {
		return err
	}

	if stock <= 0 {
		return errors.New("out of stock")
	}

	//find user cart id
	cart_id, err := i.repo.GetCartId(userID)
	if err != nil {
		return errors.New("some error in getting user cart")
	}
	//if user  has no existing cart create new cart
	if cart_id == 0 {
		cart_id, err = i.repo.CreateNewCart(userID)
		if err != nil {
			return errors.New("cannot create cart for user")
		}
	}
	exists, err := i.repo.CheckIfItemsIsAlreadyAdded(cart_id, inventoryID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("item already exist in cart")
	}
	//add product to line items
	if err := i.repo.AddLineItems(cart_id, inventoryID); err != nil {
		return errors.New("error in adding products")
	}
	return nil

}

func (i *cartUseCase) CheckOut(id int) (models.CheckOut, error) {

	address, err := i.repo.GetAddresses(id)
	if err != nil {
		return models.CheckOut{}, err
	}

	paymethods, err := i.adrepo.GetPaymentMethod()
	if err != nil {
		return models.CheckOut{}, err
	}

	products, err := i.userUsecase.GetCart(id)

	if err != nil {
		return models.CheckOut{}, err
	}

	var checkout models.CheckOut

	checkout.CartID = products.ID
	checkout.Addresses = address
	checkout.Products = products.Data
	checkout.PaymentMethods = paymethods

	return checkout, err
}
