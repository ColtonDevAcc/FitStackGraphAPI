package main

import (
	"net/http"

	"github.com/voodoostack/fitstackapi"
)

func authMiddleware(authTokenServices fitstackapi.AuthTokenService) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			token, err := authTokenServices.ParseTokenFromRequest(ctx, r)
			if err != nil {
				next.ServeHTTP(w, r)
			}

			ctx = fitstackapi.PutUserIDIntoContext(ctx, token.Sub)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
