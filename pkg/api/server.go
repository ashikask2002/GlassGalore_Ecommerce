package http

import (
	handler "GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	categoryHandler *handler.CategoryHandler,
	otpHandler *handler.OtpHandler,
	inventoryHandler *handler.InventoryHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	engine.GET("/validate-token", adminHandler.ValidateRefreshTokenAndCreateNewAccess)

	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler, inventoryHandler, orderHandler)

	return &ServerHTTP{engine: engine}

}
func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":3000")

	if err != nil {
		log.Fatal("gin engine couldn't start")
	}
}
