package credits

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TV2-Bachelorproject/server/model"
	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/db"
)

func TestCredits(t *testing.T) {
	model.Migrate()
	model.Seed()
	defer model.Reset()

	b, err := json.Marshal(CreateRequest{
		ProgramID:     1,
		CreditGroupID: 1,
		Persons:       []uint{1, 2},
	})

	if err != nil {
		t.Error(err)
	}

	body := bytes.NewBuffer(b)

	r := httptest.NewRequest("POST", "/credit", body)
	w := httptest.NewRecorder()

	Create(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", w.Code)
	}

	var credits []struct{ CreditID, PersonID uint }
	db.Table("credit_persons").Find(&credits)

	if len(credits) != 2 {
		t.Errorf("expected 2 credits to exist; got %d", len(credits))
	}
}

func TestDelete(t *testing.T) {
	model.Migrate()
	model.Seed()
	defer model.Reset()

	db.Exec("DELETE FROM credit_persons")

	person := public.Person{}.Find(1)
	credit := public.Credit{
		ProgramID:     1,
		CreditGroupID: 1,
		Persons:       []public.Person{person},
	}

	db.Create(&credit)

	b, err := json.Marshal(DeleteRequest{
		CreditID: 1,
		PersonID: 1,
	})

	if err != nil {
		t.Error(err)
	}

	body := bytes.NewBuffer(b)

	r := httptest.NewRequest("DELETE", "/credit", body)
	w := httptest.NewRecorder()

	Delete(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", w.Code)
	}

	var credits []struct{ CreditID, PersonID uint }
	db.Table("credit_persons").Find(&credits)

	if len(credits) > 0 {
		t.Errorf("expected o credits to exist; got %d", len(credits))
	}
}

func TestAccept(t *testing.T) {
	model.Migrate()
	model.Seed()
	defer model.Reset()

	db.Exec("DELETE FROM credit_persons")

	person := public.Person{}.Find(1)
	credit := public.Credit{
		ProgramID:     1,
		CreditGroupID: 1,
		Persons:       []public.Person{person},
	}

	db.Create(&credit)

	b, err := json.Marshal(AcceptRequest{
		ProgramID: 1,
	})

	if err != nil {
		t.Error(err)
	}

	body := bytes.NewBuffer(b)

	r := httptest.NewRequest("PUT", "/credit/accept", body)
	w := httptest.NewRecorder()

	Accept(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", w.Code)
	}

	var credits public.Credits
	db.Where("program_id = ?", 1).Find(&credits)

	for _, credit := range credits {
		if !credit.Accepted {
			t.Error("credit should have been accepted")
		}
	}
}
