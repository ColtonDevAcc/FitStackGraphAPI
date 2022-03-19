package fitstackapi

import "errors"

var (
	ErrNotFound             = errors.New("validation error")
	ErrValidation           = errors.New("validation error")
	ErrBadCredentials       = errors.New("email/password wrong combination")
	ErrInvalidAccessToken   = errors.New("invalid access token")
	ErrNoUserIDInContext    = errors.New("no user id in context")
	ErrGenAccessToken       = errors.New("error generating access token")
	ErrUserNotAuthenticated = errors.New("error useer is not authenticated")
)
