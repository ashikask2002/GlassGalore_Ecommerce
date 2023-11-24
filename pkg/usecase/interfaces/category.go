package interfaces

import "GlassGalore/pkg/domain"

type CategoryUseCase interface {
	AddCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(category domain.Category) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategory() ([]domain.Category, error)
}
