package handler

import (
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary User sign-up
// @Description Register a new user by providing user details in JSON format
// @Accept json
// @Produce json
// @Tags USER
// @Param user body models.UserDetails true "User details in JSON format"
// @Success 201 {object} response.Response "User successfully signed up"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or constraints not satisfied"
// @Failure 500 {object} response.Response "User could not be signed up"
// @Router /users/signup [post]
func (u *UserHandler) UserSignUp(c *gin.Context) {
	var user models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	userCreated, err := u.userUseCase.UserSignUp(user)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "user could not signed up", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusCreated, "user successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, succesRes)
}

// @Summary User login
// @Description Log in a user by providing login details in JSON format
// @Accept json
// @Produce json
// @Tags USER
// @Param user body models.UserLogin true "Login details in JSON format"
// @Success 200 {object} response.Response "User successfully logged in"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or constraints not satisfied"
// @Failure 401 {object} response.Response "User could not be logged in"
// @Router /users/login [post]
func (u *UserHandler) LoginHandler(c *gin.Context) {

	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user_details, err := u.userUseCase.LoginHandler(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user could not logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", user_details, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get user details
// @Description Retrieve details of the authenticated user
// @Accept json
// @Produce json
// @Tags USER PROFILE
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved user details"
// @Failure 400 {object} response.Response "Error in retrieving user details"
// @Router /users/profile/details [get]
func (i *UserHandler) GetUserDetails(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	details, err := i.userUseCase.GetUserDetails(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", details, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Get user addresses
// @Description Retrieve addresses of the authenticated user
// @Accept json
// @Produce json
// @Tags USER PROFILE
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved user addresses"
// @Failure 400 {object} response.Response "Error in retrieving user addresses"
// @Router /users/profile [get]
func (i *UserHandler) GetAddresses(c *gin.Context) {

	idString, _ := c.Get("id")

	addresses, err := i.userUseCase.GetAddresses(idString.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "didnt get the records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "Successfully got all records", addresses, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Add user address
// @Description Add a new address for the authenticated user
// @Accept json
// @Produce json
// @Tags USER PROFILE
// @Security BearerTokenAuth
// @Param id path int true "User ID"
// @Param address body models.AddAddress true "Address details"
// @Success 200 {object} response.Response "Successfully added the address"
// @Failure 400 {object} response.Response "Error in adding the address"
// @Router /users/profile/add [post]
func (i *UserHandler) AddAddress(c *gin.Context) {

	id, _ := c.Get("id")

	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	validator.New().Struct(address)
	if err := i.userUseCase.AddAddress(id.(int), address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not added the address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added the address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Edit user details
// @Description Edit details for the authenticated user
// @Accept json
// @Produce json
// @Tags USER PROFILE
// @Security BearerTokenAuth
// @Param id path int true "User ID"
// @Param model body models.EditDetailsResponse true "User details to be edited"
// @Success 200 {object} response.Response "Successfully edited the details"
// @Failure 400 {object} response.Response "Error in editing the details"
// @Router /users/profile [put]
func (i *UserHandler) EditDetails(c *gin.Context) {

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var model models.EditDetailsResponse

	if err := c.BindJSON(&model); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := i.userUseCase.EditDetails(id, model)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error updating the values", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Succesfully Edited the details", body, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Change user password
// @Description Change the password for the authenticated user
// @Accept json
// @Produce json
// @Tags USER PROFILE
// @Security BearerTokenAuth
// @Param id path int true "User ID"
// @Param ChangePassword body models.ChangePassword true "Password change details"
// @Success 200 {object} response.Response "Password changed successfully"
// @Failure 400 {object} response.Response "Error in changing the password"
// @Router /users/profile/security/password [put]
func (i *UserHandler) ChangePassword(c *gin.Context) {

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	var ChangePassword models.ChangePassword

	if err := c.BindJSON(&ChangePassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.ChangePassword(id, ChangePassword.OldPassword, ChangePassword.Password, ChangePassword.RePassword); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "password changed succesfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get user's shopping cart
// @Description Retrieve the products in the shopping cart for the authenticated user
// @Accept json
// @Produce json
// @Tags CART MANAGEMENT
// @Security BearerTokenAuth
// @Param id path int true "User ID"
// @Success 200 {object} response.Response "Products in the shopping cart"
// @Failure 400 {object} response.Response "Error in retrieving the shopping cart"
// @Router /users/cart/get [get]
func (i *UserHandler) GetCart(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	products, err := i.userUseCase.GetCart(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully get all products in  cart", products, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Remove a product from the shopping cart
// @Description Remove a specific product from the shopping cart for the authenticated user
// @Accept json
// @Produce json
// @Tags CART MANAGEMENT
// @Security BearerTokenAuth
// @Param cart_id query int true "Cart ID"
// @Param product_id query int true "Product ID"
// @Success 200 {object} response.Response "Successfully removed from cart"
// @Failure 400 {object} response.Response "Error in removing from cart"
// @Router /users/cart/remove [delete]
func (i *UserHandler) RemoveFromCart(c *gin.Context) {
	CartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	ProductID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.RemoveFromCart(CartID, ProductID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not remove from cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully removed from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Update product quantity in the shopping cart
// @Description Update the quantity of a specific product in the shopping cart for the authenticated user
// @Accept json
// @Produce json
// @Tags CART MANAGEMENT
// @Security BearerTokenAuth
// @Param id query int true "User ID"
// @Param product query int true "Product ID"
// @Param quantity query int true "New quantity"
// @Success 200 {object} response.Response "Successfully updated the quantity"
// @Failure 400 {object} response.Response "Error in updating the quantity"
// @Router /users/cart/update [put]
func (i *UserHandler) UpdateQuantity(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	inv, err := strconv.Atoi(c.Query("product"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	qty, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.userUseCase.UpdateQuantity(id, inv, qty); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "cannot updated the quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "success fully updated the quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
