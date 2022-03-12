package jwt

import (
	"context"
	"net/http"

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
		return fitstack
	}
}

func buildToken() fitstackapi.AuthToken {

}

func (ts *TokenService) ParseToken(ctx context.Context, payload string) (fitstackapi.AuthToken, error) {

}

func (ts *TokenService) CreateRefreshToken(ctx context.Context, user fitstackapi.User, tokenID string) (string, error) {
}

func (ts *TokenService) CreateAccessToken(ctx context.Context, user fitstackapi.User, tokenID string) (string, error) {
}
