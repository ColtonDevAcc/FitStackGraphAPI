package fitstackapi

import "errors"

var (
	ErrNotFound       = errors.New("validation error")
	ErrValidation     = errors.New("validation error")
	ErrBadCredentials = errors.New("email/password wrong combination")
)
