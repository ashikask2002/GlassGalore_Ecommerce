package interfaces

import "GlassGalore/pkg/domain"

type CategoryRepository interface {
	AddCategory(category domain.Category) (domain.Category, error)
	CheckCategory(current string) (bool, error)
	UpdateCategory(category domain.Category) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategory() ([]domain.Category, error)
	CheckCategoryExist(catname string) (bool,error)
}
