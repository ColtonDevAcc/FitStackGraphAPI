package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/voodoostack/fitstackapi/config"
	"github.com/voodoostack/fitstackapi/domain"
	"github.com/voodoostack/fitstackapi/graph"
	"github.com/voodoostack/fitstackapi/jwt"
	"github.com/voodoostack/fitstackapi/postgres"
)

func main() {
	ctx := context.Background()

	config.LoadEnv(".env")

	conf := config.New()

	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	fmt.Print("working")

	router := chi.NewRouter()

	//! REPOST
	userRepo := postgres.NewUserRepo(db)

	//! SERVICES
	authTokenServices := jwt.NewTokenService(conf)
	authService := domain.NewAuthService(userRepo, authTokenServices)

	router.Use(authMiddleware(authTokenServices))
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	router.Handle("/", playground.Handler("FitStackAPI", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					AuthService: authService,
				},
			},
		),
	))

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
