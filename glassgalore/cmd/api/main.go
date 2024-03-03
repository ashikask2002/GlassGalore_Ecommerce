package main

import (
	_ "GlassGalore/cmd/api/docs"
	config "GlassGalore/pkg/config"
	di "GlassGalore/pkg/di"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

	"log"
)

// @title Go + Gin E-Commerce API Glass Galore
// @version 1.0.0
// @description Glass Galore is the platform to buy Glasses
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:3000
// @BasePath /
// @query.collection.format multi
func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config:", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()

	}
}
