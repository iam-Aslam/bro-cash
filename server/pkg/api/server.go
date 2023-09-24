package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/routes"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/config"
)

type Server struct {
	engine *gin.Engine
	port   string
}

func NewServerHTTP(cfg config.Config, authHandler interfaces.AuthHandler) *Server {

	engine := gin.New()

	engine.Use(gin.Logger())

	routes.RegisterUserRoutes(engine.Group("/api"), authHandler)

	return &Server{
		engine: engine,
		port:   cfg.Port,
	}
}

// To start server
func (s *Server) Start() error {
	return s.engine.Run(":" + s.port)
}
