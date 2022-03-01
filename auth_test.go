package fitstackapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "NiceBro",
				Email:           "colton@me.com",
				Password:        "Password",
				ConfirmPassword: "Password",
			},
			err: nil,
		}, {
			name: "Invalid Email",
			input: RegisterInput{
				Username:        "NiceBro",
				Email:           "coltonme.com",
				Password:        "Password",
				ConfirmPassword: "Password",
			},
			err: ErrValidation,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.ValidateInput()

			if tc.err != nil {
				//!we want an error
				require.ErrorIs(t, err, tc.err)
			} else {
				//! we dont want an error
				require.NoError(t, err)
			}
		})
	}
}
