package domain

import "time"

type RefreshTokenSession struct {
	ID       uint      `gorm:"primaryKey;not null"`
	TokenID  string    `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	User     User      `gorm:"foreignKey:UserID"`
	ExpireAt time.Time `gorm:"not null"`
}
