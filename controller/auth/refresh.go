package auth

import (
	"net/http"
	"time"

	"github.com/TV2-Bachelorproject/server/model/user"
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

	claims.StandardClaims.ExpiresAt = time.Now().Add(5 * time.Minute).Unix()
	refreshedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u user.User
	db.Where("token = ?", token).First(&u)
	u.Token = refreshedToken
	db.Save(&u)

	w.Write([]byte(refreshedToken))
}
