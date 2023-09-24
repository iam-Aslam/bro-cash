package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/bro-cash/server/pkg/domain"
)

type UserRepo interface {
	SaveUser(ctx context.Context, user domain.User) (uint, error)
	FindUserByPhone(ctx context.Context, phone string) (domain.User, error)
	IsUserAlreadyExistWithThisPhone(ctx context.Context, phone string) (bool, error)
}
