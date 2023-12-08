package repository

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"

	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) *couponRepository {
	return &couponRepository{
		DB: db,
	}
}

func (i *couponRepository) AddCoupon(coupon models.Coupons) error {
	if err := i.DB.Exec("INSERT INTO coupons(coupon,discount_rate,valid) values($1,$2,$3)", coupon.Coupon, coupon.DiscountRate, coupon.Valid).Error; err != nil {
		return err
	}
	return nil
}

func (i *couponRepository) GetAllCoupons() ([]domain.Coupons, error) {
	var coupons []domain.Coupons

	err := i.DB.Table("coupons").Find(&coupons).Error
	if err != nil {
		return []domain.Coupons{}, err
	}
	return coupons, nil
}

func (i *couponRepository) MakeCouponInvalid(id int) error {
	if err := i.DB.Exec("update coupons set valid='false' where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (i *couponRepository) ReactivateCoupen(id int) error {
	if err := i.DB.Exec("update coupons set valid='true' where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
