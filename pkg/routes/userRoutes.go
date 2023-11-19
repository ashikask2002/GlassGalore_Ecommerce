package routes

import (
	"GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler) {

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

		}

	}
}
