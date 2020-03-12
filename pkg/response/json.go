package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(&v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
