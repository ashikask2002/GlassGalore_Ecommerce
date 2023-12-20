package repository

import (
	"GlassGalore/pkg/domain"
	interfaces "GlassGalore/pkg/repository/interfaces"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{DB}

}

func (p *categoryRepository) AddCategory(c domain.Category) (domain.Category, error) {

	var b string

	err := p.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category", c.Category).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}

	var categoryResponse domain.Category
	err = p.DB.Raw(` 
	SELECT 
	   p.id,
	   p.category 
	    FROM 
	       categories p
	 	WHERE
		  P.category = ?
		  `, b).Scan(&categoryResponse).Error

	if err != nil {
		return domain.Category{}, err
	}
	return categoryResponse, nil
}

func (p *categoryRepository) CheckCategory(current string) (bool, error) {
	var i int
	err := p.DB.Raw("SELECT COUNT(*) FROM categories WHERE category=?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}

	if i == 0 {
		return false, err
	}

	return true, err
}

func (p *categoryRepository) UpdateCategory(category domain.Category) (domain.Category, error) {

	// Check the database connection
	// if p.DB == nil {
	// 	return domain.Category{}, errors.New("database connection is nil")
	// }

	// update the category
	fmt.Println("xxxxxxxxxxxxx", category.Category, category.ID)
	if err := p.DB.Exec("UPDATE categories SET category = $1 WHERE id =$2", category.Category, category.ID).Error; err != nil {
		return domain.Category{}, err
	}

	//Retrieve the updated category
	var body domain.Category
	if err := p.DB.First(&body, category.ID).Error; err != nil {
		return domain.Category{}, err
	}
	return body, nil
}

func (c *categoryRepository) DeleteCategory(categoryID string) error {

	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}
	if id <= 0 {
		return errors.New("id must be positive")
	}
	result := c.DB.Exec("DELETE FROM categories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that id is exist")
	}
	return nil
}

func (c *categoryRepository) GetCategory() ([]domain.Category, error) {
	var model []domain.Category
	err := c.DB.Raw("SELECT * FROM categories").Scan(&model).Error
	if err != nil {
		return []domain.Category{}, err
	}

	return model, nil
}

func (i *categoryRepository) CheckCategoryExist(catname string) (bool, error) {
	var count int
	err := i.DB.Raw("select count(*) from categories where category = ?", catname).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
