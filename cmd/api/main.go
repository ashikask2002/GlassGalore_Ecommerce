package main

import (
	"GlassGalore/docs"
	config "GlassGalore/pkg/config"
	di "GlassGalore/pkg/di"

	"log"
)

func main() {

	//	@title						Go + Gin Mobile-Mart
	//	@description				fgdh
	//	@contact.name				API Support
	//	@securityDefinitions.apikey	BearerTokenAuth
	//	@in							header
	//	@name						Authorization
	//	@securityDefinitions.apikey	Refreshtoken
	//	@in							header
	//	@name						Refreshtoken
	//	@host						localhost:8080
	//	@BasePath					/
	//	@query.collection.format	multi
	// docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.Title = "hii"
	docs.SwaggerInfo.Host = "localhost:3000"

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
