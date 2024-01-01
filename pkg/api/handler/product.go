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

// @Summary Add product
// @Description Add a new product using JSON payload
// @Accept json
// @Produce json
// @Tags ADMIN PRODUCT MANAGEMENT
// @security BearerTokenAuth
// @Param product body models.AddProducts true "Product details in JSON format"
// @Success 200 {object} response.Response "Successfully added product"
// @Failure 400 {object} response.Response "Form file error or could not add the product"
// @Router /admin/products [post]
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

// @Summary Delete product
// @Description Delete a product by providing the product ID
// @Accept json
// @Produce json
// @Tags ADMIN PRODUCT MANAGEMENT
// @security BearerTokenAuth
// @Param id query string true "Product ID to be deleted"
// @Success 200 {object} response.Response "Successfully deleted the product"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not delete the product"
// @Router /admin/products [delete]
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

// @Summary Update product stock
// @Description Update the stock of a product by providing the product ID and new stock value
// @Accept json
// @Produce json
// @Tags ADMIN PRODUCT MANAGEMENT
// @security BearerTokenAuth
// @Param product_id body string true "Product ID to be updated"
// @Param stock body int true "New stock value for the product"
// @Success 200 {object} response.Response "Successfully updated the product stock"
// @Failure 400 {object} response.Response "Fields provided in the wrong format or could not update the product stock"
// @Router /admin/products/:id/stock [put]
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

// @Summary Edit product details
// @Description Edit the details of a product by providing the product ID and updated details
// @Accept json
// @Produce json
// @Tags ADMIN PRODUCT MANAGEMENT
// @security BearerTokenAuth
// @Param id query int true "Product ID to be edited"
// @Param details body models.EditProductDetails true "Updated details for the product"
// @Success 200 {object} response.Response "Successfully edited the product details"
// @Failure 400 {object} response.Response "Problems in the ID or fields provided in the wrong format or could not edit the product details"
// @Router /admin/products/details [put]
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

// @Summary List products for user
// @Description Get a paginated list of products for users
// @Accept json
// @Produce json
// @Tags USER PRODUCT MANAGEMENT
// @Param page query int false "Page number for pagination, default is 1"
// @Success 200 {object} response.Response "Successfully retrieved the records"
// @Failure 400 {object} response.Response "Page number not in the right format or could not retrieve records"
// @Router /users/home/product [get]
func (i *ProductHandler) ListProductForUser(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	products, err := i.ProductUseCase.ListProductForUser(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary List products for admin
// @Description Get a paginated list of products for admin
// @Accept json
// @Produce json
// @Tags ADMIN PRODUCT MANAGEMENT
// @security BearerTokenAuth
// @Param page query int false "Page number for pagination, default is 1"
// @Success 200 {object} response.Response "Successfully retrieved the records"
// @Failure 400 {object} response.Response "Page number not in the right format or could not retrieve records"
// @Router /admin/products [get]
func (i *ProductHandler) LisProductforAdmin(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := i.ProductUseCase.ListProductForUser(page)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Filter products by category
// @Description Get a list of products filtered by category ID
// @Accept json
// @Produce json
// @Tags USER PRODUCT MANAGEMENT
// @Param category_id query int true "Category ID for filtering products"
// @Success 200 {object} response.Response "Successfully retrieved the product list"
// @Failure 400 {object} response.Response "Error in conversion or cannot retrieve the product list"
// @Router /users/products/filter [get]
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

// @Summary Filter products by price range
// @Description Get a list of products within the specified price range
// @Accept json
// @Produce json
// @Tags USER PRODUCT MANAGEMENT
// @Param price query int true "Minimum price for filtering products"
// @Param pricetwo query int true "Maximum price for filtering products"
// @Success 200 {object} response.Response "Successfully retrieved the product list"
// @Failure 400 {object} response.Response "Error in conversion or cannot retrieve the product list"
// @Router /users/products/filterP [get]
func (i *ProductHandler) FilterProductsByPrice(c *gin.Context) {
	Price := c.Query("price")
	PriceInt, err := strconv.Atoi(Price)

	if err != nil {
		errirRes := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errirRes)
		return
	}

	Pricetwo := c.Query("pricetwo")
	PricetwoInt, err := strconv.Atoi(Pricetwo)

	if err != nil {
		errirRes := response.ClientResponse(http.StatusBadRequest, "error in conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errirRes)
		return
	}

	productList, err := i.ProductUseCase.FilterProductsByPrice(PriceInt, PricetwoInt)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "cannot retrieve the productlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully got all product", productList, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Search products
// @Description Get a list of products based on search criteria
// @Accept json
// @Produce json
// @Tags USER PRODUCT MANAGEMENT
// @Param page query int false "Page number for pagination (default is 1)"
// @Param limit query int false "Number of items per page (default is 5)"
// @Param search body models.Search true "Search criteria in JSON format"
// @Success 200 {object} response.Response "Successfully retrieved the product list"
// @Failure 400 {object} response.Response "Error in conversion or could not get any products"
// @Router /users/products/search [get]
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

// @Summary Add or update rating for a product
// @Description Add or update the rating for a product by providing user ID, product ID, and rating
// @Accept json
// @Produce json
// @Tags USER PRODUCT MANAGEMENT
// @security BearerTokenAuth
// @Param product_id query int true "Product ID for which rating is to be added or updated"
// @Param rating query number true "Rating to be added or updated for the product (float64)"
// @Success 200 {object} response.Response "Successfully added or updated the rating"
// @Failure 400 {object} response.Response "Error in converting ID, rating, or rating update"
// @Router /users/products/rating [post]
func (i *ProductHandler) Rating(c *gin.Context) {
	idstring, _ := c.Get("id")
	id, _ := idstring.(int)
	productId := c.Query("product_id")
	productIdInt, err := strconv.Atoi(productId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in converting id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	rating := c.Query("rating")
	ratingfloat, err := strconv.ParseFloat(rating, 64)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in converting rating", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err = i.ProductUseCase.Rating(id, productIdInt, ratingfloat)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in rating the product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfully added the rating", nil, nil)
	c.JSON(http.StatusOK, succesRes)
}

// @Summary Upload images for a product
// @Description Upload images for a product by providing the product ID and image files
// @Accept multipart/form-data
// @Produce json
// @Tags ADMIN PRODUCT MANAGEMENT
// @security BearerTokenAuth
// @Param product_id query int true "Product ID for which images are to be uploaded"
// @Param files formData file true "Images to be uploaded for the product"
// @Success 200 {object} response.Response "Successfully uploaded the images"
// @Failure 400 {object} response.Response "Error in converting ID, retrieving images from form, or updating images"
// @Router /admin/products/upload_image [post]
func (i *ProductHandler) UploadImage(c *gin.Context) {
	productid := c.Query("product_id")

	productIdInt, err := strconv.Atoi(productid)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in converting id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "retreiving images from form error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		errorRes := response.ClientResponse(http.StatusBadRequest, "no files are provided", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	for _, file := range files {
		err := i.ProductUseCase.UpdateProductImage(productIdInt, file)
		if err != nil {
			errorRes := response.ClientResponse(http.StatusBadRequest, "could not change one or more images", nil, err.Error())
			c.JSON(http.StatusBadRequest, errorRes)
			return
		}
	}

	succesRes := response.ClientResponse(http.StatusOK, "succesfully added the images", nil, nil)
	c.JSON(http.StatusOK, succesRes)

}
