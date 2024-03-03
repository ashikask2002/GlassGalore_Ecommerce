package usecase

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"errors"
)

type offerUseCase struct {
	repo interfaces.OfferRepository
}

func NewOfferUseCase(repo interfaces.OfferRepository) services.OfferUseCase {
	return &offerUseCase{
		repo: repo,
	}
}

func (i *offerUseCase) AddCategoryOffer(categoryOffer models.CategorytOfferResp) error {

	if categoryOffer.CategoryID <= 0 {
		return errors.New("id must be positive")
	}

	if categoryOffer.DiscountPrice <= 0 {
		return errors.New("discount price must be positive")
	}
	if categoryOffer.OfferName == "" {
		return errors.New("not allowed empty name")
	}
	if err := i.repo.AddCategoryOffer(categoryOffer); err != nil {

		return err
	}
	return nil
}

func (i *offerUseCase) GetCategoryOffer() ([]domain.CategoryOffer, error) {
	offers, err := i.repo.GetCategoryOffer()
	if err != nil {
		return []domain.CategoryOffer{}, err
	}
	return offers, nil
}

func (i *offerUseCase) ExpireCategoryOffer(id int) error {
	if id <= 0 {
		return errors.New("id must be positive")
	}
	if err := i.repo.ExpireCategoryOffer(id); err != nil {
		return err
	}
	return nil
}
