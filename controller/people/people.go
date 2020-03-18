package people

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/private"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/TV2-Bachelorproject/server/pkg/request"
	"github.com/TV2-Bachelorproject/server/pkg/response"
)

func List(w http.ResponseWriter, r *http.Request) {
	people := private.People{}.Find()
	response.JSON(w, people)
}

func Show(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	person := private.Person{}.Find(id)

	if person.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response.JSON(w, person)
}

func Create(w http.ResponseWriter, r *http.Request) {
	person := private.Person{}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(b, &person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if errs := person.Invalid(); len(errs) > 0 {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&person)

	response.JSON(w, person)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	person := private.Person{}.Find(id)

	if err := json.Unmarshal(b, &person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if errs := person.Invalid(); len(errs) > 0 {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	if person.ID != id {
		http.Error(w, "not possible to update the id", http.StatusBadRequest)
		return
	}

	db.Save(&person)

	response.JSON(w, person)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	person := private.Person{}.Find(id)
	db.Delete(person)

	response.JSON(w, person)
}
