package handlers

import (
    "encoding/json"
    "io"
    "net/http"

    "github.com/capt-alien/datastore-zero/internal/db"
    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
)


func PutHandler(database *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key := chi.URLParam(r, "key")
		valueBytes, err := io.ReadAll(r.Body)
		if err != nil {
			JSONError(w, "could not read request body", http.StatusBadRequest)
			return
		}

		value := string(valueBytes)
		record := db.Record{Key: key, Value: value}

		if err := database.Save(&record).Error; err != nil {
			JSONError(w, "failed to save record", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	}
}
