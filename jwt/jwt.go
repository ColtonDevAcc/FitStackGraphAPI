package jwt

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	jwtGo "github.com/lestrrat-go/jwx/jwt"
	"github.com/voodoostack/fitstackapi"
	"github.com/voodoostack/fitstackapi/config"
)

var signatureType = jwa.HS256

type TokenService struct {
	Conf *config.Config
}

func NewTokenService(conf *config.Config) *TokenService {
	return &TokenService{
		Conf: conf,
	}
}

func (ts *TokenService) ParseTokenFromRequest(ctx context.Context, r *http.Request) (fitstackapi.AuthToken, error) {
	token, err := jwtGo.ParseRequest(
		r,
		jwtGo.WithValidate(true),
		jwtGo.WithIssuer(ts.Conf.JWT.Issuer),
		jwtGo.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)),
	)
	if err != nil {
		return fitstackapi.AuthToken{}, nil
	}
	return buildToken(token), nil
}

func buildToken(token jwtGo.Token) fitstackapi.AuthToken {
	return fitstackapi.AuthToken{
		ID:  token.JwtID(),
		Sub: token.Subject(),
	}
}

func (ts *TokenService) ParseToken(ctx context.Context, payload string) (fitstackapi.AuthToken, error) {
	token, err := jwtGo.Parse(
		[]byte(payload),
		jwtGo.WithValidate(true),
		jwtGo.WithIssuer(ts.Conf.JWT.Issuer),
		jwtGo.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)),
	)
	if err != nil {
		return fitstackapi.AuthToken{}, fitstackapi.ErrInvalidAccessToken
	}

	return buildToken(token), nil
}

func (ts *TokenService) CreateRefreshToken(ctx context.Context, user fitstackapi.User, tokenID string) (string, error) {
	t := jwtGo.New()

	if err := setDefaultToken(t, user, fitstackapi.RefreshTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	if err := t.Set(jwtGo.JwtIDKey, tokenID); err != nil {
		return "", fmt.Errorf("error set jwt id: %v", err)
	}

	token, err := jwtGo.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("error sign jwt: %v", err)
	}

	return string(token), nil
}

func (ts *TokenService) CreateAccessToken(ctx context.Context, user fitstackapi.User, tokenID string) (string, error) {
	t := jwtGo.New()

	if err := setDefaultToken(t, user, fitstackapi.AccessTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	if err := t.Set(jwtGo.JwtIDKey, tokenID); err != nil {
		return "", fmt.Errorf("error set jwt id: %v", err)
	}

	token, err := jwtGo.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("error sign jwt: %v", err)
	}

	return string(token), nil
}

func setDefaultToken(t jwtGo.Token, user fitstackapi.User, lifetime time.Duration, conf *config.Config) error {
	if err := t.Set(jwtGo.SubjectKey, user.ID); err != nil {
		return fmt.Errorf("error set jwt SubjectKey: %v", err)
	}

	if err := t.Set(jwtGo.IssuerKey, conf.JWT.Issuer); err != nil {
		return fmt.Errorf("error set jwt IssuerKey: %v", err)

	}

	if err := t.Set(jwtGo.IssuedAtKey, time.Now().Unix()); err != nil {
		return fmt.Errorf("error set jwt IssuedAtKey key: %v", err)

	}

	if err := t.Set(jwtGo.ExpirationKey, time.Now().Add(lifetime).Unix()); err != nil {
		return fmt.Errorf("error set jwt ExpirationKey: %v", err)

	}
	return nil
}
