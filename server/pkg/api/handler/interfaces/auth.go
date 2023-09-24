package interfaces

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	UserSignUp(ctx *gin.Context)
}
