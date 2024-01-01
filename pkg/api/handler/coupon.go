package handler

import (
	"GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	usecase interfaces.CouponUseCase
}

func NewCouponHandler(use interfaces.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		usecase: use,
	}
}

// @Summary Create new coupon
// @Description Create a new coupon using JSON payload
// @Accept json
// @Produce json
// @Tags ADMIN COUPON MANAGEMENT
// @security BearerTokenAuth
// @Param coupon body models.Coupons true "Coupon details in JSON format"
// @Success 200 {object} response.Response "Successfully added the coupon"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not add the coupon"
// @Router /admin/coupons [post]
func (i *CouponHandler) CreateNewCoupen(c *gin.Context) {
	var coupon models.Coupons

	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fieds are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := i.usecase.AddCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coudnt added the coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully added the coupon", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Get all coupons
// @Description Retrieve a list of all coupons
// @Accept json
// @Produce json
// @Tags ADMIN COUPON MANAGEMENT
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved all coupons"
// @Failure 400 {object} response.Response "Could not get all coupons"
// @Router /admin/coupons [get]
func (i *CouponHandler) GetAllCoupons(c *gin.Context) {

	categories, err := i.usecase.GetAllCoupons()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "coudnt get all coupons", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	succesres := response.ClientResponse(http.StatusOK, "successfully got all result", categories, nil)
	c.JSON(http.StatusOK, succesres)
}

// @Summary Make coupon invalid
// @Description Mark a coupon as invalid by providing its ID
// @Accept json
// @Produce json
// @Tags ADMIN COUPON MANAGEMENT
// @security BearerTokenAuth
// @Param id query int true "Coupon ID to be marked as invalid"
// @Success 200 {object} response.Response "Successfully made the coupon invalid"
// @Failure 400 {object} response.Response "Error in converting ID or making the coupon invalid"
// @Router /admin/coupons [delete]
func (i *CouponHandler) MakeCouponInvalid(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in converting id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.usecase.MakeCouponInvalid(id); err != nil {
		errores := response.ClientResponse(http.StatusBadRequest, "error in make the coupon invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errores)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully made the coupon invalid", nil, nil)
	c.JSON(http.StatusOK, succesRes)

}

// @Summary Reactivate coupon
// @Description Reactivate a coupon by providing its ID
// @Accept json
// @Produce json
// @Tags ADMIN COUPON MANAGEMENT
// @security BearerTokenAuth
// @Param id query int true "Coupon ID to be reactivated"
// @Success 200 {object} response.Response "Successfully reactivated the coupon"
// @Failure 400 {object} response.Response "Error in converting ID or reactivating the coupon"
// @Router /admin/coupons [put]
func (i *CouponHandler) ReactivateCoupen(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in converting id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := i.usecase.ReactivateCoupen(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldnt done the reactivate coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully reactivated the coupon", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}
