package auth

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

func TestLogin(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	u, err := user.New("John Doe", "john@example.com", "123456", user.Admin)

	if err != nil {
		t.Error(err)
	}

	db.Create(&u)

	b, err := json.Marshal(Credentials{
		Email:    "john@example.com",
		Password: "123456",
	})

	if err != nil {
		t.Error(err)
	}

	body := bytes.NewBuffer(b)

	r := httptest.NewRequest("POST", "/auth/login", body)
	w := httptest.NewRecorder()

	Login(w, r)

	u = user.Find(1)

	if u.Token != w.Body.String() {
		t.Error("did not regitser token to a user correctly")
	}
}
