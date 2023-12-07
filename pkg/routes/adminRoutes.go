package routes

import (
	"GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.ProductHandler, orderHandler *handler.OrderHandler) {

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
			inventorymanagement.DELETE("", inventoryHandler.DeleteProduct)
			inventorymanagement.PUT("/details", inventoryHandler.EditProductDetails)

			inventorymanagement.PUT("/:id/stock", inventoryHandler.UpdateProduct)

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
			

		}

	}
}
