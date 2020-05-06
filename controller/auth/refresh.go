package auth

import (
	"net/http"
	"time"

	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/config"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	jwt "github.com/dgrijalva/jwt-go"
)

func Refresh(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := &Claims{}

	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().SecretKey), nil
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

	claims.StandardClaims.ExpiresAt = time.Now().Add(60 * time.Minute).Unix()
	refreshedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Get().SecretKey))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := user.Find(claims.ID)

	if u.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if u.Token != token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u.Token = refreshedToken
	db.Save(&u)

	w.Write([]byte(refreshedToken))
}
