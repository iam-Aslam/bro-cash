package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase"
	usecaseinterface "github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase/interfaces"
)

const (
	user                   = "user"
	admin                  = "admin"
	AuthorizationHeaderKey = "Authorization"
	AuthorizationType      = "Bearer"
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
		response := response.ErrorResponse("Failed to bind inputs", err)
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

		response := response.ErrorResponse("Failed to sign-up for user", err)
		ctx.JSON(statusCode, response)
		return
	}
	// generate tokens
	tokenRes, err := a.generateTokens(user, userID, ctx)
	if err != nil {
		response := response.ErrorResponse("Failed to generate tokens", err)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	AuthorizationHeaderValue := AuthorizationType + " " + tokenRes.AccessToken
	ctx.Header(AuthorizationHeaderKey, AuthorizationHeaderValue)

	response := response.SuccessResponse("Successfully user sign-up completed", tokenRes)

	ctx.JSON(http.StatusCreated, response)
}

// to generate access and refresh token for user or admin by role
func (a *authHandle) generateTokens(role string, userID uint, ctx *gin.Context) (response.TokenResponse, error) {

	accessToken, err := a.usecase.GenerateAccessToken(ctx, role, userID)
	if err != nil {
		return response.TokenResponse{}, err
	}

	refreshToken, err := a.usecase.GenerateRefreshToken(ctx, role, userID)
	if err != nil {
		return response.TokenResponse{}, err
	}

	return response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
