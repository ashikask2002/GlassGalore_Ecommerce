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

// @Summary Add item to cart
// @Description Add a product to the user's shopping cart
// @Accept json
// @Produce json
// @Tags CART MANAGEMENT
// @security BearerTokenAuth
// @Param id header int true "User ID obtained from authentication"
// @Param model body models.AddToCart true "Product details to add to the cart"
// @Success 200 {object} response.Response "Successfully added to the cart"
// @Failure 400 {object} response.Response "User ID retrieval failed, or fields provided in the wrong format, or error adding to the cart"
// @Router /users/cart/add-to-cart [post]
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
	if err := i.usecase.AddToCart(model.UserID, model.ProductID, model.Quantity); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not added the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfully added to the cart", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Process checkout
// @Description Process the checkout for the user's shopping cart
// @Accept json
// @Produce json
// @Tags CHECKOUT
// @security BearerTokenAuth
// @Param id header int true "User ID obtained from authentication"
// @Success 200 {object} response.Response "Successfully processed checkout"
// @Failure 400 {object} response.Response "User ID retrieval failed, or error processing checkout"
// @Router /cart/checkout [post]
func (i *CartHandler) CheckOut(c *gin.Context) {
	userID, _ := c.Get("id")

	products, err := i.usecase.CheckOut(userID.(int))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
