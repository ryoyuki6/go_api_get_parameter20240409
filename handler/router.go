package handler

import (
	"encoding/json"
	"net/http"
	"log"
)

func Para(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()

	log.Println("Recieved query parameters:", queryParams)

	if err := json.NewEncoder(w).Encode(queryParams); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
