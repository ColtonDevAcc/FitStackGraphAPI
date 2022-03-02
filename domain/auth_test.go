package domain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/voodoostack/fitstackapi"
	"github.com/voodoostack/fitstackapi/mocks"
)

func TestAuthServices_Register(t *testing.T) {

	validInput := fitstackapi.RegisterInput{
		Username:        "bobby",
		Email:           "email@bob.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	t.Run("can register", func(t *testing.T) {
		ctx := context.Background()

		userRepo := &mocks.UserRepo{}

		userRepo.On("GetByUsername", mock.Anything, mock.Anything).
			Return(fitstackapi.User{}, fitstackapi.ErrNotFound)

		userRepo.On("GetByEmail", mock.Anything, mock.Anything).
			Return(fitstackapi.User{}, fitstackapi.ErrNotFound)

		userRepo.On("Create", mock.Anything, mock.Anything).
			Return(fitstackapi.User{
				ID:       "123",
				Username: validInput.Username,
				Email:    validInput.Email,
			}, nil)

		services := NewAuthService(userRepo)

		res, err := services.Register(ctx, validInput)
		require.NoError(t, err)

		require.NotEmpty(t, res.AccessToken)
		require.NotEmpty(t, res.User.ID)
		require.NotEmpty(t, res.User.Email)
		require.NotEmpty(t, res.User.Username)

		userRepo.AssertExpectations(t)
	})
}
