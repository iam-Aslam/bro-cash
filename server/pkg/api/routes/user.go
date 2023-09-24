package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/interfaces"
)

func RegisterUserRoutes(api *gin.RouterGroup, authHandler interfaces.AuthHandler) {

	auth := api.Group("/auth")
	{
		signUp := auth.Group("/sign-up")
		{
			signUp.POST("/", authHandler.UserSignUp)
		}
	}
}
