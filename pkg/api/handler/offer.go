package handler

import (
	"GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OfferHandler struct {
	OfferUseCase interfaces.OfferUseCase
}

func NewOfferHandler(useCase interfaces.OfferUseCase) *OfferHandler {
	return &OfferHandler{
		OfferUseCase: useCase,
	}
}

// @Summary Add category offer
// @Description Add a new offer for a category using JSON payload
// @Accept json
// @Produce json
// @Tags ADMIN OFFER MANAGEMENT
// @security BearerTokenAuth
// @Param CategoryOffer body models.CategorytOfferResp true "Category offer details in JSON format"
// @Success 200 {object} response.Response "Successfully added the category offer"
// @Failure 400 {object} response.Response "Request fields in the wrong format or constraints not satisfied"
// @Failure 500 {object} response.Response "Error in adding the category offer"
// @Router /admin/offers [post]
func (i *OfferHandler) AddCategoryOffer(c *gin.Context) {
	var CategoryOffer models.CategorytOfferResp

	if err := c.ShouldBindJSON(&CategoryOffer); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "request fields in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err := validator.New().Struct(CategoryOffer)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	fmt.Println("xxxxxxxxxxxxxxxxx", CategoryOffer)
	err = i.OfferUseCase.AddCategoryOffer(CategoryOffer)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "error in adding the category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfully added the categoryOffer", nil, nil)
	c.JSON(http.StatusOK, succesRes)

}

// @Summary Get category offers
// @Description Retrieve all category offers
// @Accept json
// @Produce json
// @Tags ADMIN OFFER MANAGEMENT
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved category offers"
// @Failure 400 {object} response.Response "Error in getting category offers"
// @Router /admin/offers [get]
func (i *OfferHandler) GetCategoryOffer(c *gin.Context) {
	categories, err := i.OfferUseCase.GetCategoryOffer()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in getting categoy offers", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfully get all category offers", categories, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Expire category offer
// @Description Expire a category offer by providing its ID
// @Accept json
// @Produce json
// @Tags ADMIN OFFER MANAGEMENT
// @security BearerTokenAuth
// @Param id query int true "ID of the category offer to expire"
// @Success 200 {object} response.Response "Successfully expired the category offer"
// @Failure 400 {object} response.Response "Error in converting the ID or deleting the category offer"
// @Router /admin/offers [delete]
func (i *OfferHandler) ExpireCategoryOffer(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in converting the id ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err = i.OfferUseCase.ExpireCategoryOffer(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in deleting the category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully expired the offer", nil, nil)
	c.JSON(http.StatusOK, succesRes)

}
