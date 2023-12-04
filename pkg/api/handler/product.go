package handler

import (
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUseCase services.ProductUseCase
}

func NewProductHandler(usecase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		ProductUseCase: usecase,
	}
}

func (i *ProductHandler) AddProduct(c *gin.Context) {

	var product models.AddProducts
	if err := c.ShouldBindJSON(&product); err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}

	ProductResponse, err := i.ProductUseCase.AddProduct(product)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not add the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully added product", ProductResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *ProductHandler) DeleteProduct(c *gin.Context) {

	productID := c.Query("id")
	err := i.ProductUseCase.DeleteProduct(productID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successres := response.ClientResponse(http.StatusOK, "Successfully deleted the Product", nil, nil)
	c.JSON(http.StatusOK, successres)
}

func (i *ProductHandler) UpdateProduct(c *gin.Context) {
	var p models.ProductUpdate

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	a, err := i.ProductUseCase.UpdateProduct(p.Productid, p.Stock)
	if err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "could not update the Product stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated the Product stock", a, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *ProductHandler) EditProductDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "problems in the id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	var model models.EditProductDetails

	err = c.BindJSON(&model)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := i.ProductUseCase.EditProductDetails(id, model)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not edit the details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited details", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *ProductHandler) ListProductForUser(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	// id := c.MustGet("id")

	// userID, ok := id.(int)
	// if !ok {
	// 	errorRes := response.ClientResponse(http.StatusForbidden, "probem in identifying user from the context", nil, err.Error())
	// 	c.JSON(http.StatusForbidden, errorRes)
	// 	return
	// }
	products, err := i.ProductUseCase.ListProductForUser(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *ProductHandler) FilterProducts(c *gin.Context) {
	CategoryID := c.Query("category_id")
	CategoryIDInt, err := strconv.Atoi(CategoryID)

	if err != nil {
		errirRes := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errirRes)
		return
	}

	productList, err := i.ProductUseCase.FilterProducts(CategoryIDInt)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "cannot retrieve the productlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all product", productList, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *ProductHandler) SearchProducts(c *gin.Context) {
	var search models.Search
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "5")
	err := c.BindJSON(&search)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page str conversio failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "limit str conversion failed", nil, err)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
    search.Page = page
    search.Limit = limit

	productList, err := i.ProductUseCase.SearchProducts(search)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "couldnt get any products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully got all product", productList, nil)
	c.JSON(http.StatusOK, succesRes)

}
