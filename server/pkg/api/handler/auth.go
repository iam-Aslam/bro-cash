package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase"
	usecaseinterface "github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase/interfaces"
)

const (
	user  = "user"
	admin = "admin"
)

type authHandle struct {
	usecase usecaseinterface.AuthUseCase
}

func NewAuthHandler(usecase usecaseinterface.AuthUseCase) interfaces.AuthHandler {
	return &authHandle{
		usecase: usecase,
	}
}

func (a *authHandle) UserSignUp(ctx *gin.Context) {

	var body request.SignUp
	// bind user details from request
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse("failed to bind inputs", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// signUp user detail
	userID, err := a.usecase.UserSignUp(ctx, body)

	if err != nil {
		statusCode := http.StatusBadRequest
		// check error is user already exist error
		if err == usecase.ErrUserAlreadyExist {
			statusCode = http.StatusConflict
		}

		response := response.ErrorResponse("failed to sign-up for user", err)
		ctx.JSON(statusCode, response)
		return
	}

	a.setupToken(user, userID, ctx)
}

// to setup token for user or admin by role
func (a *authHandle) setupToken(role string, userID uint, ctx *gin.Context) {
	fmt.Println("success: ", role, userID)
}
