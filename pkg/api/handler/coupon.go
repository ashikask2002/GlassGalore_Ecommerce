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
