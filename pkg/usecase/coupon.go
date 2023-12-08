package usecase

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
)

type couponUseCase struct {
	repository interfaces.CouponRepository
}

func NewCouponUseCase(repo interfaces.CouponRepository) *couponUseCase {
	return &couponUseCase{
		repository: repo,
	}
}

func (i *couponUseCase) AddCoupon(coupon models.Coupons) error {
	if coupon.DiscountRate < 0 {
		return errors.New("price must be positive")
	}
	if err := i.repository.AddCoupon(coupon); err != nil {
		return err
	}
	return nil
}

func (i *couponUseCase) GetAllCoupons() ([]domain.Coupons, error) {
	coupons, err := i.repository.GetAllCoupons()
	if err != nil {
		return []domain.Coupons{}, err
	}
	return coupons, nil
}

func (i *couponUseCase) MakeCouponInvalid(id int) error {

	if id <= 0 {
		return errors.New("id need to be positive number")
	}

	if err := i.repository.MakeCouponInvalid(id); err != nil {
		return err
	}
	return nil
}

func (i *couponUseCase) ReactivateCoupen(id int) error {
	if id <= 0 {
		errors.New("id need to be positive number")
	}

	if err := i.repository.ReactivateCoupen(id); err != nil {
		return err
	}
	return nil
}
