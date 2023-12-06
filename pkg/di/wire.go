//go:build wireinject
// +build wireinject

package di

import (
	http "GlassGalore/pkg/api"
	handler "GlassGalore/pkg/api/handler"
	config "GlassGalore/pkg/config"
	db "GlassGalore/pkg/db"
	repository "GlassGalore/pkg/repository"
	usecase "GlassGalore/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,

		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
