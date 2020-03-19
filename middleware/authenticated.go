package middleware

import (
	"net/http"

	"github.com/TV2-Bachelorproject/server/controller/auth"
	"github.com/TV2-Bachelorproject/server/model/user"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secret = []byte("my secrect key...")

func Authenticated(types ...user.Type) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("token")

			if token == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims := &auth.Claims{}

			t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !t.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
