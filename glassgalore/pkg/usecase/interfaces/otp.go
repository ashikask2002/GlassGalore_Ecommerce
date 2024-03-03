package interfaces

import "GlassGalore/pkg/utils/models"

type OtpUseCase interface {
	SendOTP(phone string) error
	VerifyOTP(code models.VerifyData) (models.TokenUsers, error)
}
