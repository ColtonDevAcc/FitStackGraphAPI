package graph

import (
	"context"

	"github.com/voodoostack/fitstackapi"
)

func mapUser(u fitstackapi.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (q *queryResolver) Me(ctx context.Context) (*User, error) {
	UserID, err := fitstackapi.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fitstackapi.ErrUserNotAuthenticated
	}

	return mapUser(fitstackapi.User{
		ID: UserID,
	}), nil
}
