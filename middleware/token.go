package middleware

import (
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/private"
)

func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		service := private.Service{}.FindServiceWithToken(token)

		if service.ID == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)

	})
}
