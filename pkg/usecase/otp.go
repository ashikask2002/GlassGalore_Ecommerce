package usecase

import (
	"GlassGalore/pkg/config"
	helper_interfaces "GlassGalore/pkg/helper/interfaces"
	interfaces "GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
)

type otpUseCase struct {
	cfg           config.Config
	otpRepository interfaces.OtpRepository
	helper        helper_interfaces.Helper
}

func NewOtpUseCase(cfg config.Config, repo interfaces.OtpRepository, h helper_interfaces.Helper) services.OtpUseCase {
	return &otpUseCase{
		cfg:           cfg,
		otpRepository: repo,
		helper:        h,
	}
}

func (ot *otpUseCase) SendOTP(phone string) error {
	phoneNumber := ot.helper.PhoneValidation(phone)
	if !phoneNumber {
		return errors.New("invalid mobile Number")
	}
	ok := ot.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the uer does not exist")
	}

	ot.helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	_, err := ot.helper.TwilioSendOTP(phone, ot.cfg.SERVICESID)
	if err != nil {
		fmt.Println("errrrrr", err)
		return errors.New("error ocurred while generating OTP")

	}
	return nil

}

func (ot *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUsers, error) {

	ot.helper.TwilioSetup(ot.cfg.ACCOUNTSID, ot.cfg.AUTHTOKEN)
	err := ot.helper.TwilioVerifyOTP(ot.cfg.SERVICESID, code.Code, code.PhoneNumber)
	if err != nil {
		//thid guard clause catches the error code runs only until here
		return models.TokenUsers{}, errors.New("error while verifying")
	}

	// if user is authenticated using OTP send back user details
	userDetails, err := ot.otpRepository.UserDetailsUsingPhone(code.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString, err := ot.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	var user models.UserDetailsResponse
	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: user,
		Token: tokenString,
	}, nil

}
