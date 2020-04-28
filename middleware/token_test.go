package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

func TestValidate(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	service1 := private.Service{Name: "TV2TID", Token: "TestToken"}

	db.Create(&service1)

	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Add("token", "TestToken")
	w := httptest.NewRecorder()

	Validate(success{}).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected a status code of 200 ok; got %d", w.Code)
	}

}
