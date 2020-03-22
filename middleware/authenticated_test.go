package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/controller/auth"
	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

type success struct{}

func (success) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestAuthenticated(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	u1, _ := user.New("John Doe", "john@example.com", "123456", user.Admin)
	u2, _ := user.New("Jane Doe", "Jane@example.com", "123456", user.Producer)
	db.Create(&u1)
	db.Create(&u2)

	b, err := json.Marshal(auth.Credentials{
		Email:    "john@example.com",
		Password: "123456",
	})

	if err != nil {
		t.Error(err)
	}

	body := bytes.NewBuffer(b)

	r := httptest.NewRequest("POST", "/auth/login", body)
	w := httptest.NewRecorder()

	auth.Login(w, r)
	token := w.Body.String()

	r = httptest.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()
	r.Header.Add("token", token)

	Authenticated(user.Admin)(success{}).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected a status code of 200 ok; got %d", w.Code)
	}
}
