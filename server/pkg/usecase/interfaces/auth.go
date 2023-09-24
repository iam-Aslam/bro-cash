package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/bro-cash/server/pkg/api/handler/request"
)

type AuthUseCase interface {
	UserSignUp(ctx context.Context, signUpDetails request.SignUp) (uint, error)
}
