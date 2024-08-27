package httpserver

import (
	"log"
	"net/http"

	jwtUtil "github.com/nhan1603/CloneScc/api/internal/pkg/jwt"
)

var ErrAuthenticationFailed = &Error{Status: http.StatusUnauthorized, Code: "authentication_failed", Desc: "invalid access token"}

// AuthenticateUserMiddleware authenticates token
func AuthenticateUserMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				RespondJSON(w, ErrAuthenticationFailed)
				return
			}

			if err := jwtUtil.ValidateToken(token); err != nil {
				log.Printf("Err when validating token: %+v\n", err)
				RespondJSON(w, ErrAuthenticationFailed)
				return
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
