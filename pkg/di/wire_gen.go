package di

import (
	http "GlassGalore/pkg/api"
	"GlassGalore/pkg/api/handler"
	"GlassGalore/pkg/config"
	"GlassGalore/pkg/db"
	"GlassGalore/pkg/helper"
	"GlassGalore/pkg/repository"
	"GlassGalore/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	Helper := helper.NewHelper(cfg)

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, Helper)
	userHandler := handler.NewUserHandler(userUseCase)

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, Helper)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler)

	return serverHTTP, nil

}
