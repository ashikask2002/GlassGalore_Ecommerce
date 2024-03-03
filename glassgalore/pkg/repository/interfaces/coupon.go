package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type CouponRepository interface{
	AddCoupon(coupon models.Coupons) error
	GetAllCoupons() ([]domain.Coupons,error)
	MakeCouponInvalid(id int) error
	ReactivateCoupen(id int) error
	FindCouponPrice(id int) ( int, error)
	CheckCouponValid(id int) (bool,error)

}