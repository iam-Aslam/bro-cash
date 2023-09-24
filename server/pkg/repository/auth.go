package repository

import (
	"context"

	"github.com/nikhilnarayanan623/bro-cash/server/pkg/domain"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type authDB struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) interfaces.AuthRepo {
	return &authDB{
		db: db,
	}
}

func (a *authDB) SaveRefreshTokenSession(ctx context.Context, session domain.RefreshTokenSession) error {

	query := `INSERT INTO refresh_token_sessions (token_id, user_id, expire_at) 
	VALUES ($1, $2, $3)`

	return a.db.Exec(query, session.TokenID, session.UserID, session.ExpireAt).Error
}
