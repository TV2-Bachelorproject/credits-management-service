package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/gorilla/mux"
)

func TestListUsers(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	u1, _ := user.New("John Doe", "john@example.com", "123456", user.Admin)
	u2, _ := user.New("Jane Doe", "jane@example.com", "123456", user.Producer)
	expected := user.Users{u1, u2}

	for _, u := range expected {
		db.Create(&u)
	}

	r := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	List(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code 200; got %d", w.Code)
	}

	results := user.Users{}

	json.Unmarshal(w.Body.Bytes(), &results)

	if len(results) != 2 {
		t.Errorf("expected a list containing two users; got a list of %d", len(results))
	}

	for i, result := range results {
		if result.Name != expected[i].Name {
			t.Errorf(
				"expected a user with name %s; got %s",
				expected[i].Name, result.Name,
			)
		}
	}

}

func TestShowPerson(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	expected, _ := user.New("John Doe", "john@example.com", "123456", user.Admin)
	db.Create(&expected)

	r := httptest.NewRequest("GET", "/users/1", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	w := httptest.NewRecorder()

	Show(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code 200; got %d", w.Code)
	}

	result := user.User{}
	json.Unmarshal(w.Body.Bytes(), &result)

	if expected.Name != result.Name {
		t.Errorf(
			"expected a user with name %s; got %s",
			expected.Name, result.Name,
		)
	}
}

func TestCreatePerson(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	b, err := json.Marshal(user.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "123456",
		Type:     user.Admin,
	})

	if err != nil {
		t.Error(err)
	}

	body := bytes.NewBuffer(b)

	r := httptest.NewRequest("POST", "/users", body)
	w := httptest.NewRecorder()

	Create(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", w.Code)
	}

	result := user.User{}
	json.Unmarshal(w.Body.Bytes(), &result)
	expected := user.Find(result.ID)

	if expected.Name != result.Name {
		t.Error("the result did not match expected return value")
	}

	if expected.Password == result.Password {
		t.Error("the password was not encrypted correctly")
	}
}
