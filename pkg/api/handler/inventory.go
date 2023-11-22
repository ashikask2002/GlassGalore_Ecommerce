package handler

import (
	services "GlassGalore/pkg/usecase/interfaces"
	"GlassGalore/pkg/utils/models"
	"GlassGalore/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	InvnetoryUseCase services.InvnetoryUseCase
}

func NewInventoryHandler(usecase services.InvnetoryUseCase) *InventoryHandler {
	return &InventoryHandler{
		InvnetoryUseCase: usecase,
	}
}

func (i *InventoryHandler) AddInventory(c *gin.Context) {

	var inventory models.AddInventories
	if err := c.ShouldBindJSON(&inventory); err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}

	InventoryResponse, err := i.InvnetoryUseCase.AddInventory(inventory)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not add the Inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully added Inventory", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *InventoryHandler) DeleteInventory(c *gin.Context) {

	inventoryID := c.Query("id")
	err := i.InvnetoryUseCase.DeleteInventory(inventoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successres := response.ClientResponse(http.StatusOK, "Successfully deleted the inventory", nil, nil)
	c.JSON(http.StatusOK, successres)
}

func (i *InventoryHandler) UpdateInventory(c *gin.Context) {
	var p models.InventoryUpdate

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	a, err := i.InvnetoryUseCase.UpdateInventory(p.Productid, p.Stock)
	if err != nil {
		errorres := response.ClientResponse(http.StatusBadRequest, "could not update the inventory stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorres)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated the inventory stock", a, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) EditInventoryDetails(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "problems in the id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	var model models.EditInventoryDetails

	err = c.BindJSON(&model)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = i.InvnetoryUseCase.EditInventoryDetails(id, model)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not edit the details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited details", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) ListProductForUser(c *gin.Context) {
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
	products, err := i.InvnetoryUseCase.ListProductForUser(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
