package controller

import (
	"net/http"

	"github.com/TV2-Bachelorproject/server/model/public"
)

func Programs(w http.ResponseWriter, r *http.Request) {
	programs := []public.Program{}

	public.Programs().Find(&programs)

	for _, program := range programs {
		w.Write([]byte(program.Title + "\n"))
	}
}
