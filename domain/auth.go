package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/voodoostack/fitstackapi"
	"golang.org/x/crypto/bcrypt"
)

var (
	passwordCost = bcrypt.MinCost
)

type AuthServices struct {
	UserRepo fitstackapi.UserRepo
}

func NewAuthService(ur fitstackapi.UserRepo) *AuthServices {
	return &AuthServices{
		UserRepo: ur,
	}
}

func (as *AuthServices) Register(ctx context.Context, input fitstackapi.RegisterInput) (fitstackapi.AuthResponse, error) {
	input.Sanitize()

	if err := input.ValidateInput(); err != nil {
		return fitstackapi.AuthResponse{}, err
	}

	//! does user exist ?
	if _, err := as.UserRepo.GetByUsername(ctx, input.Username); !errors.Is(err, fitstackapi.ErrNotFound) {
		return fitstackapi.AuthResponse{}, fitstackapi.ErrUserNameTaken
	}

	//! is user email taken ?
	if _, err := as.UserRepo.GetByEmail(ctx, input.Email); !errors.Is(err, fitstackapi.ErrNotFound) {
		return fitstackapi.AuthResponse{}, fitstackapi.ErrEmailTaken
	}

	user := fitstackapi.User{
		Email:    input.Email,
		Username: input.Username,
	}

	//! hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), passwordCost)
	if err != nil {
		return fitstackapi.AuthResponse{}, fmt.Errorf("error hasing password: %w", err)
	}
	user.Password = string(hashPassword)

	//! create the user
	user, err = as.UserRepo.Create(ctx, user)
	if err != nil {
		return fitstackapi.AuthResponse{}, fmt.Errorf("error creating user: %v", err)
	}

	return fitstackapi.AuthResponse{
		AccessToken: "token",
		User:        user,
	}, nil
}

func (as *AuthServices) Login(ctx context.Context, input fitstackapi.LoginInput) (fitstackapi.AuthResponse, error) {
	input.Sanitize()

	if err := input.ValidateInput(); err != nil {
		return fitstackapi.AuthResponse{}, err
	}

	user, err := as.UserRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		switch {
		case errors.Is(err, fitstackapi.ErrNotFound):
			return fitstackapi.AuthResponse{}, fitstackapi.ErrBadCredentials

		default:
			return fitstackapi.AuthResponse{}, err

		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return fitstackapi.AuthResponse{}, fitstackapi.ErrBadCredentials
	}

	return fitstackapi.AuthResponse{
		AccessToken: "token",
		User:        user,
	}, nil
}
