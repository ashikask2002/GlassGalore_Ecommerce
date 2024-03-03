package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type OfferRepository interface {
	AddCategoryOffer(categoryoffer models.CategorytOfferResp) error
	GetCategoryOffer() ([]domain.CategoryOffer, error)
	ExpireCategoryOffer(id int) error
}
