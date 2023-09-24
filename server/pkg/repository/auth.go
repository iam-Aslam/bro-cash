package repository

import (
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
