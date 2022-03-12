package graph

import (
	"context"
	"errors"

	"github.com/voodoostack/fitstackapi"
)

func mapAuthResponse(a fitstackapi.AuthResponse) *AuthResponse {
	return &AuthResponse{
		AccessToken: a.AccessToken,
		User:        mapUser(a.User),
	}
}

func (m *mutationResolver) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	res, err := m.AuthService.Register(ctx, fitstackapi.RegisterInput{
		Email:           input.Email,
		Username:        input.Username,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	})
	if err != nil {
		switch {
		case errors.Is(err, fitstackapi.ErrNotFound) ||
			errors.Is(err, fitstackapi.ErrEmailTaken) ||
			errors.Is(err, fitstackapi.ErrUserNameTaken):
			return nil, buildBadrequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
}
func (m *mutationResolver) LoginInput(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	res, err := m.AuthService.Login(ctx, fitstackapi.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, fitstackapi.ErrNotFound) ||
			errors.Is(err, fitstackapi.ErrEmailTaken) ||
			errors.Is(err, fitstackapi.ErrUserNameTaken):
			return nil, buildBadrequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil

}
