package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/nikhilnarayanan623/bro-cash/server/pkg/config"
)

type jwtAuth struct {
	key string
}

// New TokenAuth for jwt based token auth
func NewJwtTokenAuth(cfg config.Config) TokenAuth {

	return &jwtAuth{
		key: cfg.JwtKey,
	}
}

var (
	ErrInvalidExpireTime  = errors.New("payload expire time is already elapsed")
	ErrInvalidToken       = errors.New("invalid token")
	ErrFailedToParseToken = errors.New("failed to parse token to claims")
	ErrExpiredToken       = errors.New("token expired")
)

type jwtClaims struct {
	UserID    interface{}
	Email     string
	TokenID   string
	ExpiresAt time.Time
	Role      string
}

// Generate JWT token for the given payload
func (c *jwtAuth) GenerateToken(payload Payload) (TokenResponse, error) {

	// validate the given expire time is valid
	if time.Since(payload.ExpireAt) > 0 {
		return TokenResponse{}, ErrInvalidExpireTime
	}
	// create a jwt claims with given payload
	claims := &jwtClaims{
		TokenID:   payload.TokenID,
		UserID:    payload.UserID,
		Email:     payload.Email,
		Role:      payload.Role,
		ExpiresAt: payload.ExpireAt,
	}
	// if token id not provided by user create a random id for it
	if claims.TokenID == "" {
		claims.TokenID = generateTokenID()
	}
	// create a new token with claims
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with key
	tokenString, err := tkn.SignedString([]byte(c.key))

	if err != nil {
		return TokenResponse{}, fmt.Errorf("failed to sign the token: %w", err)
	}

	response := TokenResponse{
		TokenID:     claims.TokenID,
		TokenString: tokenString,
	}

	return response, nil
}

// Verify JWT token string and return payload
func (c *jwtAuth) VerifyToken(tokenString string) (Payload, error) {
	// parse the token string to jwt claims
	tkn, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(c.key), nil
	})

	if err != nil {
		// check is token expired or any other
		if errors.Is(err, ErrExpiredToken) {
			return Payload{}, ErrExpiredToken
		}
		return Payload{}, ErrInvalidToken
	}

	// parse the token into claims
	claims, ok := tkn.Claims.(*jwtClaims)
	if !ok {
		return Payload{}, ErrFailedToParseToken
	}

	response := Payload{
		TokenID:  claims.TokenID,
		UserID:   claims.UserID,
		Email:    claims.Email,
		Role:     claims.Role,
		ExpireAt: claims.ExpiresAt,
	}
	return response, nil
}

// Validate claims
func (c *jwtClaims) Valid() error {
	if time.Since(c.ExpiresAt) > 0 {
		return ErrExpiredToken
	}
	return nil
}

// Generate a random unique token id
func generateTokenID() string {
	return uuid.New().String()
}
