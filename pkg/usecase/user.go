package usecase

import (
	"GlassGalore/pkg/domain"
	helper_interface "GlassGalore/pkg/helper/interfaces"
	interfaces "GlassGalore/pkg/repository/interfaces"
	use "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"

	"errors"
)

type UserUseCase struct {
	userRepo interfaces.UserRepository
	helper   helper_interface.Helper
}

func NewUserUseCase(repo interfaces.UserRepository, helper helper_interface.Helper) use.UserUseCase {
	return &UserUseCase{
		userRepo: repo,
		helper:   helper,
	}
}

var InternalError = "Internal server Error"
var ErrorHashingPassword = "Error In Hashing Password"

func (u *UserUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {
	number := u.helper.PhoneValidation(user.Phone)
	if !number {
		return models.TokenUsers{}, errors.New("invalid mobilenumber")
	}

	email := u.helper.IsValidEmail(user.Email)
	if !email {
		return models.TokenUsers{}, errors.New("invalid email")
	}

	userExist := u.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")

	}
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	hashedPassword, err := u.helper.PasswordHashing(user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New(ErrorHashingPassword)
	}
	user.Password = hashedPassword

	userData, err := u.userRepo.UserSignUp(user)

	if err != nil {
		return models.TokenUsers{}, errors.New("could not add the user")
	}

	tokenString, err := u.helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	return models.TokenUsers{
		Users: userData,
		Token: tokenString,
	}, nil

}

func (u *UserUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")

	}
	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, errors.New(InternalError)
	}
	if isBlocked {
		return models.TokenUsers{}, errors.New("user is blocked by admin")
	}

	// Get the user details in order to check the password, in this case (The same function can be reused in future)
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, errors.New(InternalError)
	}

	err = u.helper.CompareHashAndPassword(user_details.Password, user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse

	userDetails.Id = int(user_details.Id)
	userDetails.Name = user_details.Name
	userDetails.Email = user_details.Email
	userDetails.Phone = user_details.Phone

	tokenString, err := u.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")

	}
	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

func (i *UserUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {
	details, err := i.userRepo.GetUserDetails(id)
	if err != nil {
		return models.UserDetailsResponse{}, errors.New("error in getting details")
	}

	return details, nil
}

func (i *UserUseCase) GetAddresses(id int) ([]domain.Address, error) {
	addresses, err := i.userRepo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return addresses, nil
}

func (i *UserUseCase) AddAddress(id int, address models.AddAddress) error {

	phone := i.helper.PhoneValidation(address.Phone)
	if !phone {
		return errors.New("invalid phone Number")
	}

	pin := i.helper.IsValidPIN(address.Pin)
	if !pin {
		return errors.New("invalid pin number")
	}

	rslt := i.userRepo.CheckIfFirstAddress(id)
	var result bool

	if !rslt {
		result = true
	} else {
		result = false
	}

	err := i.userRepo.AddAddress(id, address, result)
	if err != nil {
		return errors.New("error in adding address")
	}
	return nil

}

func (i *UserUseCase) EditDetails(id int, user models.EditDetailsResponse) (models.EditDetailsResponse, error) {
	if !i.helper.PhoneValidation(user.Phone) {
		return models.EditDetailsResponse{}, errors.New("phone number is invalid")
	}

	if !i.helper.IsValidEmail(user.Email) {
		return models.EditDetailsResponse{}, errors.New("email is invalid")
	}

	body, err := i.userRepo.EditDetails(id, user)
	if err != nil {
		return models.EditDetailsResponse{}, err
	}

	return body, nil
}

// func (i *UserUseCase) EditEmail(id int, email string) error {

// 	err := i.userRepo.EditEmail(id, email)
// 	if err != nil {
// 		return errors.New("could not change email")
// 	}
// 	return nil
// }

// func (i *UserUseCase) EditPhone(id int, Phone string) error {
// 	err := i.userRepo.EditPhone(id, Phone)
// 	if err != nil {
// 		return errors.New("could not change phone")
// 	}
// 	return nil
// }

func (i *UserUseCase) ChangePassword(id int, old string, password string, repassword string) error {
	userPassword, err := i.userRepo.GetPassword(id)
	if err != nil {
		return errors.New(InternalError)
	}

	err = i.helper.CompareHashAndPassword(userPassword, old)
	if err != nil {
		return errors.New("password incorrect")
	}

	if password != repassword {
		return errors.New("password does not match")
	}
	newPassword, err := i.helper.PasswordHashing(password)
	if err != nil {
		return errors.New("error in Hashing password")
	}

	return i.userRepo.ChangePassword(id, string(newPassword))
}

func (u *UserUseCase) GetCart(id int) (models.GetCartResponse, error) {
	//find cart id

	cart_id, err := u.userRepo.GetCartID(id)

	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}
	//find product inside the cart
	products, err := u.userRepo.GetProductsInCart(cart_id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}
	//find products names
	var product_names []string
	for i := range products {
		product_name, err := u.userRepo.FindProductNames(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		product_names = append(product_names, product_name)
	}

	//find quantity
	var quantity []int
	for i := range products {
		q, err := u.userRepo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		quantity = append(quantity, q)
	}

	//find price
	var price []float64
	for i := range products {
		q, err := u.userRepo.FindPrice(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}

		price = append(price, q)
	}

	var stocks []int
	for i := range products {
		stock, err := u.userRepo.FindStock(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		stocks = append(stocks, stock)
	}

	//find quantity
	var categories []int
	for i := range products {
		c, err := u.userRepo.FindCategory(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}

		categories = append(categories, c)
	}

	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart
		get.ID = products[i]
		get.ProductName = product_names[i]
		get.Category_id = categories[i]
		get.Quantity = quantity[i]
		get.Total = price[i] * float64(quantity[i])
		get.Price = int(price[i])
		get.StockAvailabe = stocks[i]
		get.DiscountedPrice = 0

		getcart = append(getcart, get)
	}

	var offers []float64

	for i := range categories {

		c, err := u.userRepo.GetCatOfferr(categories[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		offers = append(offers, c)
	}

	//find discounted price
	for i := range getcart {
		getcart[i].DiscountedPrice = (getcart[i].Total) - offers[i]
	}

	var response models.GetCartResponse
	response.ID = cart_id
	response.Data = getcart

	//return in appropriate format
	return response, nil
}

func (i *UserUseCase) RemoveFromCart(cart, inventory int) error {
	if cart <= 0 {
		return errors.New("cartid must be positive")
	}

	if inventory <= 0 {
		return errors.New("inventoryid must be positive")
	}
	err := i.userRepo.RemoveFromCart(cart, inventory)
	if err != nil {
		return err
	}
	return nil
}

func (i *UserUseCase) UpdateQuantity(id, inv_id, qty int) error {
	if id <= 0 {
		return errors.New(" id must be positive")
	}

	if inv_id <= 0 {
		return errors.New("inventory id must be positive")
	}
	if qty <= 0 {
		return errors.New("quantity must be positive")
	}
	err := i.userRepo.UpdateQuantity(id, inv_id, qty)
	if err != nil {
		return err
	}
	return nil
}
