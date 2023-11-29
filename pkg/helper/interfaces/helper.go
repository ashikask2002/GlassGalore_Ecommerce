package interfaces

import "GlassGalore/pkg/utils/models"

type Helper interface {
	PasswordHashing(string) (string, error)
	GenerateTokenClients(user models.UserDetailsResponse) (string, error)
	CompareHashAndPassword(a string, b string) error
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error)
	TwilioSetup(username string, password string)
	TwilioSendOTP(phone string, serviceID string) (string, error)
	TwilioVerifyOTP(serviceID string, code string, phone string) error
	PhoneValidation(phone string) bool
}
