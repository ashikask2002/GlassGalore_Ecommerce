package interfaces

import (
	"GlassGalore/pkg/domain"
	"GlassGalore/pkg/utils/models"
)

type OfferUseCase interface {
	AddCategoryOffer(CategoryOffer models.CategorytOfferResp) error
	GetCategoryOffer()([]domain.CategoryOffer,error)
	ExpireCategoryOffer(id int) error
}
