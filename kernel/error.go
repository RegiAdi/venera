package kernel

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUserUpdateFailed       = errors.New("user update failed")
	ErrPasswordUnmatch        = errors.New("password do not match")
	ErrGenerateAPITokenFailed = errors.New("failed to generate api token")
	ErrInvalidObjectID        = errors.New("invalid objectid")
	ErrUserAlreadyExists      = errors.New("username already exists")
	ErrEmailAlreadyExists     = errors.New("email already exists")

	// Product related errors
	ErrProductNotFound     = errors.New("product not found")
	ErrProductCreateFailed = errors.New("failed to create product")
	ErrProductUpdateFailed = errors.New("failed to update product")
	ErrProductDeleteFailed = errors.New("failed to delete product")
)
