package handler

import (
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"errors"
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

// @Summary Order items from cart
// @Description Create an order by providing order details including address, payment method, and coupon
// @Accept json
// @Produce json
// @Tags CHECKOUT
// @security BearerTokenAuth
// @Param id query int true "User ID"
// @Param order body models.Order true "Order details in JSON format"
// @Success 200 {object} response.Response "Successfully made the order"
// @Failure 400 {object} response.Response "Error in getting user ID, fields provided in the wrong format, or could not make the order"
// @Router /users/check-out/order [post]
func (i *OrderHandler) OrderItemsFromCart(c *gin.Context) {
	userID, ok := c.Get("id")
	userIDint := userID.(int)
	if !ok {
		err := errors.New("didnt got id")
		erroRes := response.ClientResponse(http.StatusBadRequest, "error in getting id", nil, err.Error())
		c.JSON(http.StatusBadRequest, erroRes)
	}

	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong foramt", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	order.UserID = userIDint

	if err := i.orderUseCase.OrderItemsFromCart(order.UserID, order.AddressID, order.PaymentMethodID, order.CouponID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not make the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully made the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Get orders
// @Description Retrieve orders based on the provided order ID
// @Accept json
// @Produce json
// @Tags USER ORDER MANAGEMENT
// @security BearerTokenAuth
// @Param order_id query int true "Order ID to retrieve specific order, omit for all orders"
// @Success 200 {object} response.Response "Successfully retrieved orders"
// @Failure 400 {object} response.Response "Check your order ID again or could not get the orders"
// @Router /users/profile/orders [get]
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

// @Summary Get all orders
// @Description Retrieve all orders for the authenticated user with optional pagination
// @Accept json
// @Produce json
// @Tags USER ORDER MANAGEMENT
// @security BearerTokenAuth
// @Param page query int false "Page number for pagination, default is 1"
// @Param count query int false "Number of orders per page, default is 10"
// @Success 200 {object} response.Response "Successfully retrieved all orders"
// @Failure 400 {object} response.Response "Page number or count not in the correct format or could not retrieve orders"
// @Router /users/profile/orders/all [get]
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

// @Summary Cancel order
// @Description Cancel a specific order by providing the order ID
// @Accept json
// @Produce json
// @Tags USER ORDER MANAGEMENT
// @security BearerTokenAuth
// @Param order_id query int true "ID of the order to cancel"
// @Success 200 {object} response.Response "Successfully cancelled the order"
// @Failure 400 {object} response.Response "Check your order ID again or could not cancel the order"
// @Router /users/profile/orders [delete]
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

// @Summary Get admin orders
// @Description Retrieve a paginated list of orders for administrative purposes
// @Accept json
// @Produce json
// @Tags ADMIN ORDER MANAGEMENT
// @security BearerTokenAuth
// @Param page query int true "Page number for pagination"
// @Success 200 {object} response.Response "Successfully retrieved admin orders"
// @Failure 400 {object} response.Response "Page number not in the right format or could not retrieve orders"
// @Router /admin/orders [get]
func (i *OrderHandler) GetAdminOrders(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	orders, err := i.orderUseCase.GetAdminOrders(page)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve the orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all orders", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Approve order
// @Description Approve a specific order by providing its ID
// @Accept json
// @Produce json
// @Tags ADMIN ORDER MANAGEMENT
// @security BearerTokenAuth
// @Param order_id query string true "ID of the order to be approved"
// @Success 200 {object} response.Response "Successfully approved the order"
// @Failure 400 {object} response.Response "Could not approve the order or incorrect order ID format"
// @Router /admin/orders [patch]
func (i *OrderHandler) ApproveOrder(c *gin.Context) {
	orderID := c.Query("order_id")

	err := i.orderUseCase.OrdersStatus(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not approve order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succeRes := response.ClientResponse(http.StatusOK, "successfully approved order", nil, nil)
	c.JSON(http.StatusOK, succeRes)
}

// @Summary Return order
// @Description Initiate a return process for a specific order by providing its ID
// @Accept json
// @Produce json
// @Tags USER ORDER MANAGEMENT
// @security BearerTokenAuth
// @Param order_id query string true "ID of the order to be returned"
// @Success 200 {object} response.Response "Successfully initiated the return process"
// @Failure 400 {object} response.Response "Error in converting order ID or returning order"
// @Router /users/profile/orders [put]
func (i *OrderHandler) ReturnOrder(c *gin.Context) {
	id := c.Query("order_id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = i.orderUseCase.ReturnOrder(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in returning order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully returned the order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Print invoice
// @Description Generate and download the invoice for a specific order by providing its ID
// @Accept json
// @Produce json
// @Tags CHECKOUT
// @security BearerTokenAuth
// @Param order_id query string true "ID of the order for which the invoice is to be printed"
// @Success 200 {object} response.Response "Successfully downloaded the invoice"
// @Failure 400 {object} response.Response "Error in converting order ID or printing invoice"
// @Failure 502 {object} response.Response "Error in printing invoice"
// @Router /users/check-out/invoice [get]
func (i *OrderHandler) PrintInvoice(c *gin.Context) {
	orderId := c.Query("order_id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in converting id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	pdf, err := i.orderUseCase.PrintInvoice(orderIdInt)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err.Error())
		c.JSON(http.StatusBadGateway, errorRes)
		return
	}

	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")

	err = pdf.Output(c.Writer)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err.Error())
		c.JSON(http.StatusBadGateway, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully downloaded the pdf", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}
