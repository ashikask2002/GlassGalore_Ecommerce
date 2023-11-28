package routes

import (
	"GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler) {

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

		inventorymanagement := engine.Group("/inventories")
		{
			inventorymanagement.POST("", inventoryHandler.AddInventory)
			inventorymanagement.DELETE("", inventoryHandler.DeleteInventory)
			inventorymanagement.PUT("/details", inventoryHandler.EditInventoryDetails)

			inventorymanagement.PUT("/:id/stock", inventoryHandler.UpdateInventory)

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
