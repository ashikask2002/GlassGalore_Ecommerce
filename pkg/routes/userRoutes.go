package routes

import (
	"GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler,
	inventoryHandler *handler.InventoryHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)
	engine.Use(middleware.UserAuthMiddleware)

	{
		profile := engine.Group("/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("", userHandler.GetAddresses)
			profile.POST("/add", userHandler.AddAddress)

			edit := profile.Group("/edit")
			{
				edit.PUT("/name", userHandler.EditName)
				edit.PUT("/email", userHandler.EditEmail)
				edit.PUT("/phone", userHandler.EditPhone)
			}

			secutiry := profile.Group("/security")
			{
				secutiry.PUT("password", userHandler.ChangePassword)
			}

			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.GET("/all", orderHandler.GetAllOrders)
				orders.DELETE("", orderHandler.CancelOrder)
			}

		}

		// cart := engine.Group("/cart")
		// {
		// 	cart.GET("/", userHandler.GetCart)
		// }

		home := engine.Group("/home")
		{
			home.GET("/product", inventoryHandler.ListProductForUser)
		}

		cart := engine.Group("/cart")
		{
			cart.POST("add-to-cart", cartHandler.AddToCart)
			cart.GET("get", userHandler.GetCart)
			cart.DELETE("/remove", userHandler.RemoveFromCart)
			cart.PUT("update", userHandler.UpdateQuantity)
		}
		checkout := engine.Group("/check-out")
		{
			checkout.GET("", cartHandler.CheckOut)
			checkout.POST("/order", orderHandler.OrderItemsFromCart)
		}

	}
}
