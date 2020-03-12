package controller

import (
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/request"
	"github.com/TV2-Bachelorproject/server/pkg/response"
)

func People(w http.ResponseWriter, r *http.Request) {
	people := public.People{}.Find()
	response.JSON(w, people)
}

func Person(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	person := public.Person{}.Find(uint(id))
	response.JSON(w, person)
}
