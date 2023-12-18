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

}
