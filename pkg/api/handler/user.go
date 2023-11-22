package handler

import (
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"fmt"
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

func (i *UserHandler) GetUserDetails(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	fmt.Println("zzzz", id)

	details, err := i.userUseCase.GetUserDetails(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", details, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) GetAddresses(c *gin.Context) {

	idString, _ := c.Get("id")
	// id, err := strconv.Atoi(idString)
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "Check yout id again", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }

	addresses, err := i.userUseCase.GetAddresses(idString.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "didnt get the records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "Successfully got all records", addresses, nil)
	c.JSON(http.StatusOK, succesRes)
}

func (i *UserHandler) AddAddress(c *gin.Context) {

	id, _ := c.Get("id")
	// id, err := strconv.Atoi(c.Query("id"))
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }

	var address models.AddAddress
	if err := c.BindJSON(&address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.AddAddress(id.(int), address); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not added the address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added the address", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) EditName(c *gin.Context) {

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	// id, err := strconv.Atoi(c.Query("id"))

	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }

	var model models.EditName

	if err := c.BindJSON(&model); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := i.userUseCase.EditName(id, model.Name); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not edit the name", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "Succesfully changed the name", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) EditEmail(c *gin.Context) {

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	// id, err := strconv.Atoi(c.Query("id"))
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }

	var model models.EditEmail
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.EditEmail(id, model.Email); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldnt edit the email", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "email successfully edited", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) EditPhone(c *gin.Context) {
	idString, _ := c.Get("id")
	id, _ := idString.(int)

	// id, err := strconv.Atoi(c.Query("id"))
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }

	var model models.EditPhone

	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.EditPhone(id, model.Phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldnt edit  the phone", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "phone edited successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) ChangePassword(c *gin.Context) {

	idString, _ := c.Get("id")
	id, _ := idString.(int)

	// id, err := strconv.Atoi(c.Query("id"))
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }

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

func (i *UserHandler) RemoveFromCart(c *gin.Context) {
	CartID, err := strconv.Atoi(c.Query("cart_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	InventoryID, err := strconv.Atoi(c.Query("inventory_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.userUseCase.RemoveFromCart(CartID, InventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not remove from cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully removed from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *UserHandler) UpdateQuantity(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	inv, err := strconv.Atoi(c.Query("inventory"))
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
