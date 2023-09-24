package usecase

import "errors"

var (
	ErrUserAlreadyExist = errors.New("user already exist with given details")
)
