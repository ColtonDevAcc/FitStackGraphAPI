package fitstackapi

import (
	"context"
	"time"
)

var (
	AccessTokenLifetime  = time.Minute * 15   //! 15 min
	RefreshTokenLifetime = time.Hour * 24 * 7 //! 1 week
)

type RefreshToken struct {
	ID         string
	Name       string
	UserID     string
	LastUsedAt time.Time
	ExpiredAt  time.Time
	CreatedAt  time.Time
}

type CreateRefreshToken struct {
	Sub  string
	Name string
}

type RefreshTokenRepo interface {
	Create(ctx context.Context, params CreateRefreshToken) (*RefreshToken, error)
	GetByID(ctx context.Context, id string) (RefreshToken, error)
}
