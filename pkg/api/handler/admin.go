package handler

import (
	"GlassGalore/pkg/helper"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// @Summary Admin login
// @Description Authenticate admin user and generate access token
// @Accept json
// @Produce json
// @Tags ADMIN
// @Param adminDetails body models.AdminLogin true "Admin login details in JSON format"
// @Success 200 {object} response.Response "Admin authenticated successfully"
// @Failure 400 {object} response.Response "Invalid request format or authentication failure"
// @Router /admin/adminlogin [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) {

	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in corect format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.Set("Access", admin.AccessToken)
	c.Set("Refresh", admin.RefreshToken)

	successRes := response.ClientResponse(http.StatusOK, "admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Get users
// @Description Retrieve a paginated list of users
// @Accept json
// @Produce json
// @Tags ADMIN
// @security BearerTokenAuth
// @Param page query int true "Page number for pagination"
// @Success 200 {object} response.Response "Successfully retrieved the users"
// @Failure 400 {object} response.Response "Page number not in the right format or could not retrieve records"
// @Router /admin/users [get]
func (ad *AdminHandler) Getusers(c *gin.Context) {

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ad.adminUseCase.GetUsers(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Block user
// @Description Block a user by their ID
// @Accept json
// @Produce json
// @Tags ADMIN
// @security BearerTokenAuth
// @Param id query string true "User ID to be blocked"
// @Success 200 {object} response.Response "Successfully blocked the user"
// @Failure 400 {object} response.Response "User could not be blocked"
// @Router /admin/users/block [post]
func (ad *AdminHandler) BlockUser(c *gin.Context) {

	id := c.Query("id")
	err := ad.adminUseCase.BlockUser(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Unblock user
// @Description Unblock a user by their ID
// @Accept json
// @Produce json
// @Tags ADMIN
// @security BearerTokenAuth
// @Param id query string true "User ID to be unblocked"
// @Success 200 {object} response.Response "Successfully unblocked the user"
// @Failure 400 {object} response.Response "User could not be unblocked"
// @Router /admin/users/unblock [post]
func (ad *AdminHandler) UnBlockUser(c *gin.Context) {
	id := c.Query("id")
	err := ad.adminUseCase.UnBlockUser(id)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user could not be unblocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (a *AdminHandler) ValidateRefreshTokenAndCreateNewAccess(c *gin.Context) {

	refreshToken := c.Request.Header.Get("RefreshToken")

	// Check if the refresh token is valid.
	_, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("refreshsecret"), nil
	})
	if err != nil {
		// The refresh token is invalid.
		c.AbortWithError(401, errors.New("refresh token is invalid:user have to login again"))
		return
	}

	claims := &helper.AuthcustomClaims{
		Role: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newAccessToken, err := token.SignedString([]byte("accesssecret"))
	if err != nil {
		c.AbortWithError(500, errors.New("error in creating new access token"))
	}

	c.JSON(200, newAccessToken)
}

// @Summary Add new payment method
// @Description Add a new payment method using JSON payload
// @Accept json
// @Produce json
// @Tags ADMIN PAYMENT MANAGEMENT
// @security BearerTokenAuth
// @Param method body models.NewPaymentMethod true "New payment method details in JSON format"
// @Success 200 {object} response.Response "Successfully added payment method"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not add the payment method"
// @Router /admin/payment-method [post]
func (i *AdminHandler) NewPaymentMethod(c *gin.Context) {
	var method models.NewPaymentMethod
	if err := c.BindJSON(&method); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := i.adminUseCase.NewPaymentMethod(method.PaymenMethod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not added the payment method", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succcesRes := response.ClientResponse(http.StatusOK, "successfully added payment method", nil, nil)
	c.JSON(http.StatusOK, succcesRes)

}

// @Summary List payment methods
// @Description Retrieve a list of all available payment methods
// @Accept json
// @Produce json
// @Tags ADMIN PAYMENT MANAGEMENT
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved all payment methods"
// @Failure 400 {object} response.Response "Cannot list the payment methods"
// @Router /admin/payment-method [get]
func (i *AdminHandler) ListPaymentMethods(c *gin.Context) {
	categories, err := i.adminUseCase.ListPaymentMethods()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "cannot list the payment methods", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusOK, "successfully got all payment method", categories, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Delete payment method
// @Description Delete a payment method by its ID
// @Accept json
// @Produce json
// @Tags ADMIN PAYMENT MANAGEMENT
// @security BearerTokenAuth
// @Param id query int true "Payment method ID to be deleted"
// @Success 200 {object} response.Response "Successfully deleted the payment method"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or error in deleting data"
// @Router /admin/payment-method [delete]
func (i *AdminHandler) DeletePaymentMethod(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = i.adminUseCase.DeletePaymentMethod(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in delete data", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully deleted the category", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get dashboard details
// @Description Retrieve details for the admin dashboard
// @Accept json
// @Produce json
// @Tags ADMIN DASHBOARD
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved dashboard details"
// @Failure 400 {object} response.Response "Error in getting dashboard details"
// @Router /admin/dashboard [get]
func (i *AdminHandler) DashBoard(c *gin.Context) {
	dashboard, err := i.adminUseCase.DashBoard()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in getting dashboard details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	succesRes := response.ClientResponse(http.StatusOK, "successfully got all details of the  dashboard", dashboard, nil)
	c.JSON(http.StatusOK, succesRes)
}

// FilteredSalesReport retrieves the  current sales report for a specified time period.
// @Summary Retrieve current sales report for a specific time period
// @Description Retrieves sales report for the specified time period
// @Tags ADMIN DASHBOARD
// @Accept json
// @Produce json
// @security BearerTokenAuth
// @Param period query string true "Time period for sales report"
// @Success 200 {object} response.Response "Sales report retrieved successfully"
// @Failure 500 {object} response.Response "Unable to retrieve sales report"
// @Router /admin/salesreport  [get]
func (i *AdminHandler) Salesreport(c *gin.Context) {
	timePeriod := c.Query("period")

	salesReport, err := i.adminUseCase.FilteredSalesReport(timePeriod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in getting sales report", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	c.Header("Content-Disposition", "attachment; filename=sales_report.pdf")
	c.Header("Content-Type", "application/pdf")

	err = salesReport.Output(c.Writer)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadGateway, "error in printing invoice", nil, err.Error())
		c.JSON(http.StatusBadGateway, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully got the sales report", salesReport, nil)
	c.JSON(http.StatusOK, succesRes)
}
