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

func (Cat *CategoryHandler) UpdateCategory(c *gin.Context) {

	var category domain.Category

	// id_str := c.Param("id")
	// id, err := strconv.Atoi(id_str)

	// if err != nil {
	// 	errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
	// 	fmt.Println("wwwwwwwwwwwwww", id)
	// 	c.JSON(http.StatusBadRequest, errRes)
	// 	return
	// }

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
