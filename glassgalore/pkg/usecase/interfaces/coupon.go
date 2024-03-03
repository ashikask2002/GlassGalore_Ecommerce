package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type CouponUseCase interface {
	AddCoupon(coupon models.Coupons) error
	GetAllCoupons() ([]domain.Coupons, error)
	MakeCouponInvalid(id int) error
	ReactivateCoupen(id int) error
}
