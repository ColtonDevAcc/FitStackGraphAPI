package fitstackapi

import (
	"context"
	"errors"
	"time"
)

var (
	ErrUserNameTaken = errors.New("username taken")
	ErrEmailTaken    = errors.New("email taken")
)

type UserRepo interface {
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
