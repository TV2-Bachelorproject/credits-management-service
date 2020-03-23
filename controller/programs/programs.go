package programs

import (
	"fmt"
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/public"
	"github.com/TV2-Bachelorproject/server/pkg/request"
	"github.com/TV2-Bachelorproject/server/pkg/response"
)

//GetAll returns []programs
func GetAll(w http.ResponseWriter, r *http.Request) {
	programs := public.Programs{}.Find()
	response.JSON(w, programs)
}

//Get returns a program
func Get(w http.ResponseWriter, r *http.Request) {
	id, err := request.GetID(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	program := public.Program{}.Find(id)

	if program.ID == 0 {
		http.Error(w, "Program not found", http.StatusNotFound)
		fmt.Println(program)
		return
	}

	response.JSON(w, program)
}
