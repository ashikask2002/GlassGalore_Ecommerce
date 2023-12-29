package routes

import (
	"GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, offerHandler *handler.OfferHandler) {

	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.GET("", adminHandler.Getusers)
			usermanagement.PUT("/block", adminHandler.BlockUser)
			usermanagement.PUT("/unblock", adminHandler.UnBlockUser)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.POST("", categoryHandler.AddCategory)
			categorymanagement.PATCH("", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("", categoryHandler.DeleteCategory)
			categorymanagement.GET("", categoryHandler.GetCategory)
		}

		inventorymanagement := engine.Group("/products")
		{
			inventorymanagement.POST("", inventoryHandler.AddProduct)
			inventorymanagement.GET("", inventoryHandler.LisProductforAdmin)

			inventorymanagement.DELETE("", inventoryHandler.DeleteProduct)
			inventorymanagement.PUT("/details", inventoryHandler.EditProductDetails)

			inventorymanagement.PUT("/:id/stock", inventoryHandler.UpdateProduct)
			inventorymanagement.POST("/upload_image", inventoryHandler.UploadImage)

		}

		payment := engine.Group("/payment-method")
		{
			payment.POST("", adminHandler.NewPaymentMethod)
			payment.GET("", adminHandler.ListPaymentMethods)
			payment.DELETE("", adminHandler.DeletePaymentMethod)
		}

		orders := engine.Group("/orders")
		{
			orders.GET("", orderHandler.GetAdminOrders)
			orders.PATCH("", orderHandler.ApproveOrder)
			// orders.PUT("", orderHandler.ReturnOrder)

		}

		coupen := engine.Group("/coupons")
		{
			coupen.POST("", couponHandler.CreateNewCoupen)
			coupen.GET("", couponHandler.GetAllCoupons)
			coupen.DELETE("", couponHandler.MakeCouponInvalid)
			coupen.PUT("", couponHandler.ReactivateCoupen)
		}

		engine.GET("/dashboard", adminHandler.DashBoard)
		engine.GET("/salesreport", adminHandler.Salesreport)

		offer := engine.Group("/offers")
		{
			offer.POST("", offerHandler.AddCategoryOffer)
			offer.GET("", offerHandler.GetCategoryOffer)
			offer.DELETE("", offerHandler.ExpireCategoryOffer)
		}

	}
}
