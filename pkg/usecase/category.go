package usecase

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
	"errors"
)

type categoryUseCase struct {
	repository interfaces.CategoryRepository
}

func NewCategoryUseCase(repo interfaces.CategoryRepository) services.CategoryUseCase {
	return &categoryUseCase{
		repository: repo,
	}
}

func (Cat *categoryUseCase) AddCategory(category domain.Category) (domain.Category, error) {
	if category.Category == "" {
		return domain.Category{}, errors.New("category not be empty")
	}
	productResponse, err := Cat.repository.AddCategory(category)

	if err != nil {
		return domain.Category{}, err
	}

	return productResponse, nil
}

func (Cat *categoryUseCase) UpdateCategory(category domain.Category) (domain.Category, error) {
    if category.ID <= 0{
		return domain.Category{},errors.New("id must be positive")
	}
	if category.Category == "" {
		return domain.Category{}, errors.New("must provide a value")
	}

	newcat, err := Cat.repository.UpdateCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return newcat, nil
}

func (Cat *categoryUseCase) DeleteCategory(categoryID string) error {
	err := Cat.repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil
}

func (Cat *categoryUseCase) GetCategory() ([]domain.Category, error) {
	categories, err := Cat.repository.GetCategory()

	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}
