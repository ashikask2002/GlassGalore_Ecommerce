package handler

import (
	"GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase interfaces.CartUseCase
}

func NewCartHandler(usecase interfaces.CartUseCase) *CartHandler {
	return &CartHandler{
		usecase: usecase,
	}
}

func (i *CartHandler) AddToCart(c *gin.Context) {
	var model models.AddToCart
	UserID, err := c.Get("id")
	if !err {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, errors.New("getting user Id is failed"))
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	model.UserID = UserID.(int)
	if err := i.usecase.AddToCart(model.UserID, model.InventoryID, model.Quantity); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not added the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfully added to the cart", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}

func (i *CartHandler) CheckOut(c *gin.Context) {
	userID, _ := c.Get("id")
	// if err != nil {
	// 	errorRes := response.ClientResponse(http.StatusBadRequest, "user_id not in correct format", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errorRes)
	// 	return
	// }
	products, err := i.usecase.CheckOut(userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
