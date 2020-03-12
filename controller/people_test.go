package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

func TestPeople(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	expected := public.Person{Name: "John"}
	db.Create(&expected)

	r, err := http.NewRequest("GET", "/people", nil)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	People(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code 200; got %d", w.Code)
	}

	result := []public.Person{}

	json.Unmarshal(w.Body.Bytes(), &result)

	if len(result) != 1 {
		t.Errorf("expected a list containing a single person; got a list of %d", len(result))
	}

	if result[0].Name != expected.Name {
		t.Errorf(
			"expected a person with name %s; got %s",
			expected.Name, result[0].Name,
		)
	}
}

// func TestPerson(t *testing.T) {
// 	model.Migrate()
// 	defer model.Reset()
//
// 	expected := public.Person{Name: "John"}
// 	db.Create(&expected)
//
// 	r, err := http.NewRequest("GET", "/people/1", nil)
// 	context.Set(r, 0, map[string]string{
// 		"id": "1",
// 	})
// 	r.WithContext(context)
//
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	w := httptest.NewRecorder()
//
// 	Person(w, r)
//
// 	if w.Code != http.StatusOK {
// 		t.Errorf("expected status code 200; got %d", w.Code)
// 		t.Log(w.Body.String())
// 	}
//
// }
