package repository

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/repository/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type OfferRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(DB *gorm.DB) interfaces.OfferRepository {
	return &OfferRepository{
		DB: DB,
	}
}

func (i *OfferRepository) AddCategoryOffer(categoryoffer models.CategorytOfferResp) error {
	fmt.Println("nnnnnnnnnnnnnnnnn", categoryoffer)
	var count int
	err := i.DB.Raw("select count(*) from category_offers where offer_name = ? and category_id = ?", categoryoffer.OfferName, categoryoffer.CategoryID).Scan(&count).Error
	if err != nil {
		return errors.New("error in adding details")
	}
	fmt.Println("count is ", count)
	if count >= 1 {
		return errors.New("the offer is already exist")
	}

	err = i.DB.Raw("select count(*) from category_offers where category_id= ?", categoryoffer.CategoryID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		err = i.DB.Exec("delete from category_offers where category_id = ?", categoryoffer.CategoryID).Error
		if err != nil {
			return err
		}
	}

	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 5)
	err = i.DB.Exec("insert into category_offers (category_id,offer_name,discount_price, start_date,end_date) values(?,?,?,?,?)", categoryoffer.CategoryID, categoryoffer.OfferName, categoryoffer.DiscountPrice, startDate, endDate).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *OfferRepository) GetCategoryOffer() ([]domain.CategoryOffer, error) {
	var categoryOfferDetails []domain.CategoryOffer
	err := i.DB.Raw("select * from category_offers").Scan(&categoryOfferDetails).Error
	if err != nil {
		return []domain.CategoryOffer{}, errors.New("error in getting categoryoffer details")
	}
	return categoryOfferDetails, nil
}

func (i *OfferRepository) ExpireCategoryOffer(id int) error {
	if err := i.DB.Exec("delete from category_offers where id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
