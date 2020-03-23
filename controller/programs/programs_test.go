package programs

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/gorilla/mux"
)

func TestGetAll(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	expected := []public.Program{
		public.Program{Title: "Westworld"},
		public.Program{Title: "Westworld 2"},
	}

	for i := range expected {
		db.Create(&expected[i])
	}

	r := httptest.NewRequest("GET", "/programs", nil)
	w := httptest.NewRecorder()

	GetAll(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200; got %d", w.Code)
	}

	results := []public.Program{}

	json.Unmarshal(w.Body.Bytes(), &results)

	if len(results) != 2 {
		t.Errorf("Expected a list of 2 programs; got a list with %d", len(results))
	}

	for i, result := range results {
		if result.Title != expected[i].Title {
			t.Errorf("Expected a program title which was %s; but got %s",
				expected[i].Title,
				result.Title)
		}
	}

}

func TestGet(t *testing.T) {
	model.Migrate()
	defer model.Reset()

	expected := public.Program{Title: "Westworld 2"}
	db.Create(&expected)

	r := httptest.NewRequest("GET", "/programs/1", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	w := httptest.NewRecorder()

	Get(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200; got %d", w.Code)
	}

	result := public.Program{}

	json.Unmarshal(w.Body.Bytes(), &result)

	if result.Title != expected.Title {
		t.Errorf("Expected a program title which was %s; but got %s",
			expected.Title,
			result.Title)
	}

}
