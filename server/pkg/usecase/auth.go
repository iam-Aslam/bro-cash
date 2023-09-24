package usecase

import (
	"context"

	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/domain"
	repo "github.com/nikhilnarayanan623/bro-cash/server/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/utils"
)

type authUseCase struct {
	authRepo repo.AuthRepo
	userRepo repo.UserRepo
}

func NewAuthUseCase(ar repo.AuthRepo, ur repo.UserRepo) interfaces.AuthUseCase {
	return &authUseCase{
		authRepo: ar,
		userRepo: ur,
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
