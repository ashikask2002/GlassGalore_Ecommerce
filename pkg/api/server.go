package http

import (
	handler "GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	categoryHandler *handler.CategoryHandler,
	otpHandler *handler.OtpHandler,
	inventoryHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymendHandler,
	walletHandler *handler.WalletHandler,
	couponHandler *handler.CouponHandler,
	offerHandler *handler.OfferHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.LoadHTMLGlob("template/*")

	// use ginSwagger middleware to serve the API docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.GET("/validate-token", adminHandler.ValidateRefreshTokenAndCreateNewAccess)

	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler, walletHandler, couponHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler, inventoryHandler, orderHandler, couponHandler, offerHandler)

	return &ServerHTTP{engine: engine}

}
func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":3000")

	if err != nil {
		log.Fatal("gin engine couldn't start")
	}
}
