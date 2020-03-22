package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/user"
	"github.com/TV2-Bachelorproject/server/pkg/db"
	"github.com/TV2-Bachelorproject/server/pkg/request"
	"github.com/TV2-Bachelorproject/server/pkg/response"
)

func List(w http.ResponseWriter, r *http.Request) {
	users := user.All()
	response.JSON(w, users)
}

func Show(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u := user.Find(id)

	if u.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response.JSON(w, u)
}

func Create(w http.ResponseWriter, r *http.Request) {
	u := user.User{}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(b, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	u, err = user.New(u.Name, u.Email, u.Password, u.Type)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if errs := u.Invalid(); len(errs) > 0 {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&u)

	response.JSON(w, u)
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

	user := user.Find(id)

	if err := json.Unmarshal(b, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if errs := user.Invalid(); len(errs) > 0 {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	if user.ID != id {
		http.Error(w, "not possible to update the id", http.StatusBadRequest)
		return
	}

	db.Save(&user)

	response.JSON(w, user)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user := user.Find(id)
	db.Delete(user)

	response.JSON(w, user)
}
