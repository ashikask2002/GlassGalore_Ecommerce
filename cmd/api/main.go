package main

import (

	config "GlassGalore/pkg/config"
	di "GlassGalore/pkg/di"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "GlassGalore/cmd/api/docs"

	"log"
)

// @title Go + Gin E-Commerce API Watch Hive
// @version 1.0.0
// @description Watch Hive is an E-commerce platform to purchase Watch
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8080
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
