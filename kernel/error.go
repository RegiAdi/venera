package kernel

import "errors"

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrUserUpdateFailed       = errors.New("user update failed")
	ErrPasswordUnmatch        = errors.New("password do not match")
	ErrGenerateAPITokenFailed = errors.New("failed to generate api token")
	ErrInvalidObjectID        = errors.New("invalid objectid")
)
