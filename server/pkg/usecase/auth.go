package usecase

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/domain"
	repo "github.com/nikhilnarayanan623/bro-cash/server/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/token"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/utils"
)

const (
	AccessTokenDuration  = 20 * time.Minute
	RefreshTokenDuration = 24 * 7 * time.Hour
)

type authUseCase struct {
	authRepo  repo.AuthRepo
	userRepo  repo.UserRepo
	tokenAuth token.TokenAuth
}

func NewAuthUseCase(ar repo.AuthRepo, ur repo.UserRepo, tokenAuth token.TokenAuth) interfaces.AuthUseCase {
	return &authUseCase{
		authRepo:  ar,
		userRepo:  ur,
		tokenAuth: tokenAuth,
	}
}

func (a *authUseCase) UserSignUp(ctx context.Context, signUpDetails request.SignUp) (uint, error) {

	// first check the user already exist or not
	userExist, err := a.userRepo.IsUserAlreadyExistWithThisPhone(ctx, signUpDetails.Phone)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to check user already exist in db")
	}
	if userExist {
		return 0, ErrUserAlreadyExist
	}

	// hash user password
	hashPass, err := utils.GenerateHashFromPassword(signUpDetails.Password)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to hash user password")
	}

	user := domain.User{
		FirstName: signUpDetails.FirstName,
		LastName:  signUpDetails.LastName,
		Phone:     signUpDetails.Phone,
		Password:  hashPass,
	}
	// save user details on database
	userID, err := a.userRepo.SaveUser(ctx, user)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to save user details on db")
	}

	return userID, nil
}

func (a *authUseCase) GenerateAccessToken(ctx context.Context, role string, userID uint) (string, error) {

	var (
		tokenID  = utils.GenerateUniqueString()
		expireAt = time.Now().Add(AccessTokenDuration)
	)
	payload := token.Payload{
		TokenID:  tokenID,
		UserID:   userID,
		Role:     role,
		ExpireAt: expireAt,
	}

	tokenRes, err := a.tokenAuth.GenerateToken(payload)
	if err != nil {
		return "", utils.PrependMessageToError(err, "failed to generate access token")
	}

	return tokenRes.TokenString, nil
}

func (a *authUseCase) GenerateRefreshToken(ctx context.Context, role string, userID uint) (string, error) {

	var (
		tokenID  = utils.GenerateUniqueString()
		expireAt = time.Now().Add(RefreshTokenDuration)
	)
	payload := token.Payload{
		TokenID:  tokenID,
		UserID:   userID,
		Role:     role,
		ExpireAt: expireAt,
	}
	tokenRes, err := a.tokenAuth.GenerateToken(payload)
	if err != nil {
		return "", utils.PrependMessageToError(err, "failed to generate refresh token")
	}

	// store the refresh token details on db
	refreshSession := domain.RefreshTokenSession{
		TokenID:  tokenID,
		UserID:   userID,
		ExpireAt: expireAt,
	}
	if err = a.authRepo.SaveRefreshTokenSession(ctx, refreshSession); err != nil {
		return "", utils.PrependMessageToError(err, "failed to save refresh token details on db")
	}

	return tokenRes.TokenString, nil
}
