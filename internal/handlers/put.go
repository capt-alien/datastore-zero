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

        id := chi.URLParam(r, "id")
        if id == "" {
            JSONError(w, "missing id", http.StatusBadRequest)
            return
        }

        body, err := io.ReadAll(r.Body)
        if err != nil || len(body) == 0 {
            JSONError(w, "could not read request body", http.StatusBadRequest)
            return
        }

        record := db.Record{ID: id, Value: string(body)}
        if err := database.Create(&record).Error; err != nil {
            JSONError(w, "failed to save record", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
    }
}
