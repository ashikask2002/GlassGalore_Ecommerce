package handler

import (
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {

	return &OrderHandler{
		orderUseCase: useCase,
	}
}

func (i *OrderHandler) OrderItemsFromCart(c *gin.Context) {

	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong foramt", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.orderUseCase.OrderItemsFromCart(order.UserID, order.AddressID, order.PaymentMethodID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully made the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *OrderHandler) GetOrders(c *gin.Context) {
	idString := c.Query("order_id")
	order_id, err := strconv.Atoi(idString)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orders, err := i.orderUseCase.GetOrders(order_id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldnt get the orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all orders", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *OrderHandler) GetAllOrders(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number correct in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("count", "10"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, _ := c.Get("id")
	UserID, _ := id.(int)

	orders, err := i.orderUseCase.GerAllOrders(UserID, page, pageSize)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all orders", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *OrderHandler) CancelOrder(c *gin.Context) {
	idString := c.Query("order_id")
	orderID, err := strconv.Atoi(idString)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = i.orderUseCase.CancelOrder(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not cancel the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully cancelled the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
