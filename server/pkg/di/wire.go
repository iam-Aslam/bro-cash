//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/nikhilnarayanan623/bro-cash/server/pkg/api"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/config"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/db"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/repository"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {

	wire.Build(
		db.ConnectToDatabase,
		repository.NewAuthRepo,
		repository.NewUserRepo,
		usecase.NewAuthUseCase,
		handler.NewAuthHandler,
		http.NewServerHTTP,
	)

	return &http.Server{}, nil
}
