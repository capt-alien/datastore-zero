package handlers

import (
    "encoding/json"
    "net/http"
)

func HireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := map[string]string{
		"message": "Hey ASE, this took less than 24 hours to build. Imagine what I’ll do with a badge and a paycheck.",
	}
	json.NewEncoder(w).Encode(message)
}
