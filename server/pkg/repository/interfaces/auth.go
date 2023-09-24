package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/bro-cash/server/pkg/domain"
)

type AuthRepo interface {
	SaveRefreshTokenSession(ctx context.Context, session domain.RefreshTokenSession) error
}
