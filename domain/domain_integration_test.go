package domain

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/voodoostack/fitstackapi"
	"github.com/voodoostack/fitstackapi/config"
	"github.com/voodoostack/fitstackapi/jwt"
	"github.com/voodoostack/fitstackapi/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	conf             *config.Config
	db               *postgres.DB
	authTokenService fitstackapi.AuthTokenService
	authService      fitstackapi.AuthService
	userRepo         fitstackapi.UserRepo
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	config.LoadEnv(".env.test")

	passwordCost = bcrypt.MinCost

	conf = config.New()

	db = postgres.New(ctx, conf)
	defer db.Close()

	if err := db.Drop(); err != nil {
		log.Fatal(err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	userRepo = postgres.NewUserRepo(db)

	authTokenService = jwt.NewTokenService(conf)

	authService = NewAuthService(userRepo, authTokenService)

	os.Exit(m.Run())
}
