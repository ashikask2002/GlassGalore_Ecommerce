package handler

import (
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: useCase,
	}
}

// @Summary Send OTP
// @Description Send a One-Time Password (OTP) to the specified phone number
// @Accept json
// @Produce json
// @Tags USER
// @Param phone body models.OTPData true "Phone number details in JSON format"
// @Success 200 {object} response.Response "OTP sent successfully"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not send OTP"
// @Router /users/otplogin [post]
func (ot *OtpHandler) SendOTP(c *gin.Context) {

	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
	}

	err := ot.otpUseCase.SendOTP(phone.PhoneNumber)
	if err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "could not send otp", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "OTP send successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Verify OTP
// @Description Verify the provided One-Time Password (OTP) code
// @Accept json
// @Produce json
// @Tags USER
// @Param code body models.VerifyData true "Verification code details in JSON format"
// @Success 200 {object} response.Response "Successfully verified OTP"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not verify OTP"
// @Router /users/verifyotp [post]
func (ot *OtpHandler) VerifyOTP(c *gin.Context) {

	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	users, err := ot.otpUseCase.VerifyOTP(code)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not verify the OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully verified OTP", users, nil)
	c.JSON(http.StatusOK, successRes)

}
