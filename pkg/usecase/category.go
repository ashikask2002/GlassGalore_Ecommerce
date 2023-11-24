package usecase

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	services "GlassGalore/pkg/usecase/interfaces"
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
	productResponse, err := Cat.repository.AddCategory(category)

	if err != nil {
		return domain.Category{}, err
	}

	return productResponse, nil
}

func (Cat *categoryUseCase) UpdateCategory(category domain.Category) (domain.Category, error) {

	// result, err := Cat.repository.CheckCategory(current)
	// if err != nil {
	// 	return domain.Category{}, err
	// }
	// if !result {
	// 	return domain.Category{}, errors.New("there is no category as you mentioned")
	// }

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
