// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/config"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/db"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/repository"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	gormDB, err := db.ConnectToDatabase(cfg)
	if err != nil {
		return nil, err
	}
	authRepo := repository.NewAuthRepo(gormDB)
	userRepo := repository.NewUserRepo(gormDB)
	authUseCase := usecase.NewAuthUseCase(authRepo, userRepo)
	authHandler := handler.NewAuthHandler(authUseCase)
	server := api.NewServerHTTP(cfg, authHandler)
	return server, nil
}
