package people

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/gorilla/mux"
)

func TestListPeople(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	expected := []public.Person{
		public.Person{Name: "John Doe"},
		public.Person{Name: "Jane Doe"},
	}

	for i := range expected {
		db.Create(&expected[i])
	}

	r, err := http.NewRequest("GET", "/people", nil)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	List(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code 200; got %d", w.Code)
	}

	results := []public.Person{}

	json.Unmarshal(w.Body.Bytes(), &results)

	if len(results) != 2 {
		t.Errorf("expected a list containing two people; got a list of %d", len(results))
	}

	for i, result := range results {
		if result.Name != expected[i].Name {
			t.Errorf(
				"expected a person with name %s; got %s",
				expected[i].Name, result.Name,
			)
		}
	}

}

func TestShowPerson(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	expected := public.Person{Name: "John Doe"}
	db.Create(&expected)

	r := httptest.NewRequest("GET", "/people/1", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	w := httptest.NewRecorder()

	Show(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code 200; got %d", w.Code)
	}

	result := public.Person{}
	json.Unmarshal(w.Body.Bytes(), &result)

	if expected.Name != result.Name {
		t.Errorf(
			"expected a person with name %s; got %s",
			expected.Name, result.Name,
		)
	}
}

func TestCreatePerson(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	b, err := json.Marshal(private.Person{
		Name:    "John Doe",
		Email:   "john@example.com",
		Address: "Østre Stationsvej 100",
		City:    "Odense",
		Postal:  "5000",
		Country: "Denmark",
	})

	if err != nil {
		t.Error(err)
	}

	body := bytes.NewBuffer(b)

	r := httptest.NewRequest("POST", "/people", body)
	w := httptest.NewRecorder()

	Create(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", w.Code)
	}

	result := private.Person{}
	json.Unmarshal(w.Body.Bytes(), &result)
	expected := private.Person{}.Find(result.ID)

	if expected.Name != result.Name {
		t.Error("the result did not match expected return value")
	}
}
