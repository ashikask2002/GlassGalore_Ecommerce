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

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository, Helper)
	otpHandler := handler.NewOtpHandler(otpUseCase)

	inventoryRepository := repository.NewInventoryRepository(gormDB)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository, Helper)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)

	cartRepository := repository.NewCartRepository(gormDB)
	cartUsecase := usecase.NewCartUseCase(cartRepository, inventoryRepository)
	cartHandler := handler.NewCartHandler(cartUsecase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, categoryHandler, otpHandler, inventoryHandler, cartHandler)

	return serverHTTP, nil

}
