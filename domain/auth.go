package domain

import "github.com/voodoostack/fitstackapi"

type AuthServices struct {
	UserRepo fitstackapi.UserRepo
}

func NewAuthService(userRepo fitstackapi.UserRepo) *AuthServices {
	return &AuthServices{
		UserRepo: userRepo,
	}
}
