package middleware

import (
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/model/user"
)

func Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonToken := r.Header.Get("token")

		if jsonToken != "" {
			Authenticated(user.Admin, user.Producer)(next).ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("service-token")

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
