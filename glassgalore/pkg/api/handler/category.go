package handler

import (
	"GlassGalore/pkg/domain"
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUseCase services.CategoryUseCase
}

func NewCategoryHandler(usecase services.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		CategoryUseCase: usecase,
	}
}

// @Summary Add category
// @Description Add a new category using JSON payload
// @Accept json
// @Produce json
// @Tags CATEGORY MANAGEMENT
// @security BearerTokenAuth
// @Param category body domain.Category true "Category details in JSON format"
// @Success 200 {object} response.Response "Successfully added category"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not add the category"
// @Router /admin/category [post]
func (Cat *CategoryHandler) AddCategory(c *gin.Context) {
	var category domain.Category
	if err := c.BindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	CategoryResponse, err := Cat.CategoryUseCase.AddCategory(category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not add the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Catefory", CategoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Update category
// @Description Update an existing category using JSON payload
// @Accept json
// @Produce json
// @Tags CATEGORY MANAGEMENT
// @security BearerTokenAuth
// @Param category body domain.Category true "Category details in JSON format"
// @Success 200 {object} response.Response "Successfully updated category"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not update the category"
// @Router /admin/category [patch]
func (Cat *CategoryHandler) UpdateCategory(c *gin.Context) {

	var category domain.Category

	if err := c.BindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	body, err := Cat.CategoryUseCase.UpdateCategory(category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully renamed the category", body, nil)
	c.JSON(http.StatusOK, successRes)
}
// @Summary Delete category
// @Description Delete an existing category by ID
// @Accept json
// @Produce json
// @Tags CATEGORY MANAGEMENT
// @security BearerTokenAuth
// @Param id query string true "Category ID to be deleted"
// @Success 200 {object} response.Response "Successfully deleted category"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not delete the category"
// @Router /admin/category [delete]
func (Cat *CategoryHandler) DeleteCategory(c *gin.Context) {

	categoryID := c.Query("id")
	err := Cat.CategoryUseCase.DeleteCategory(categoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the category", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
// @Summary Get categories
// @Description Retrieve a list of all categories
// @Accept json
// @Produce json
// @Tags CATEGORY MANAGEMENT
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved all categories"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not retrieve categories"
// @Router /admin/category [get]
func (Cat *CategoryHandler) GetCategory(c *gin.Context) {

	categories, err := Cat.CategoryUseCase.GetCategory()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	SuccessRes := response.ClientResponse(http.StatusOK, "Successfully got all Categories", categories, nil)
	c.JSON(http.StatusOK, SuccessRes)
}
