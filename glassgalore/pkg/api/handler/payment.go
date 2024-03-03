package handler

import (
	"GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymendHandler struct {
	paymentusecase interfaces.PaymentUseCase
}

func NewPaymentHandler(usecase interfaces.PaymentUseCase) *PaymendHandler {
	return &PaymendHandler{
		paymentusecase: usecase,
	}
}
// @Summary Make payment using RazorPay
// @Description Initiate the payment process using RazorPay for the specified user and order
// @Accept json
// @Produce html
// @Tags USER
// @Param user_id query string true "User ID for whom the payment is being made"
// @Param order_id query string true "Order ID for which the payment is being made"
// @Success 200 {string} html "HTML page for initiating the RazorPay payment"
// @Failure 400 {object} response.Response "Error in the payment process"
// @Router /users/payment [get]
func (i *PaymendHandler) MakePaymentRazorPay(c *gin.Context) {
	userID := c.Query("user_id")

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "there is error check your user id again", nil, errors.New("errors in convert userId into string"+err.Error()))
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orderID := c.Query("order_id")
	orderIDint, err := strconv.Atoi(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "ther is an error check your order id again", nil, errors.New("errors in convert orderId into string"+err.Error()))
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fmt.Println("orderId handler", orderID)
	body, razorID, err := i.paymentusecase.MakePaymentRazorPay(userIDint, orderIDint)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error happened", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	// fmt.Println("body now%$#$%#$", body)
	// fmt.Println("body next", body.FinalPrice, razorID, userID, body.OrderID, body.Name, body.FinalPrice)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"final_price": body.FinalPrice * 100,
		"razor_id":    razorID,
		"user_id":     userID,
		"order_id":    body.OrderID,
		"user_name":   body.Name,
		"total":       int(body.FinalPrice),
	})
}

func (i *PaymendHandler) VerifyPayment(c *gin.Context) {
	orderID := c.Query("order_id")
	paymentID := c.Query("payment_id")
	razorID := c.Query("razor_id")

	if err := i.paymentusecase.SavePaymentDetails(paymentID, razorID, orderID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}
