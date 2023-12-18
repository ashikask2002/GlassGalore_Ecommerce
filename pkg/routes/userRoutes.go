package routes

import (
	"GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymendHandler,
	WallerHandler *handler.WalletHandler,
	couponHandler *handler.CouponHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)
	engine.GET("/payment", paymentHandler.MakePaymentRazorPay)
	engine.GET("/verifypayment", paymentHandler.VerifyPayment)

	engine.Use(middleware.UserAuthMiddleware)

	{
		profile := engine.Group("/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("", userHandler.GetAddresses)
			profile.POST("/add", userHandler.AddAddress)
			profile.PUT("", userHandler.EditDetails)

			secutiry := profile.Group("/security")
			{
				secutiry.PUT("password", userHandler.ChangePassword)
			}

			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.GET("/all", orderHandler.GetAllOrders)
				orders.DELETE("", orderHandler.CancelOrder)
				orders.PUT("", orderHandler.ReturnOrder)
			}

		}

		// cart := engine.Group("/cart")
		// {
		// 	cart.GET("/", userHandler.GetCart)
		// }

		home := engine.Group("/home")
		{
			home.GET("/product", productHandler.ListProductForUser)
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

		products := engine.Group("/products")
		{
			products.POST("/search", productHandler.SearchProducts)
			products.POST("/filter", productHandler.FilterProducts)
			products.POST("filterP",productHandler.FilterProductsByPrice)
		}

		wallet := engine.Group("wallet")
		{
			wallet.GET("", WallerHandler.ViewWallet)
		}

		coupon := engine.Group("coupon")
		{
			coupon.GET("",couponHandler.GetAllCoupons)
			
		}

	}
}
