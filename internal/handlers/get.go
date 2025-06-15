package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/capt-alien/datastore-zero/internal/db"
    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
)

func GetHandler(database *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        id := chi.URLParam(r, "id")

        var record db.Record
        err := database.First(&record, "id = ?", id).Error
        if err != nil {
            JSONError(w, "record not found", http.StatusNotFound)
            return
        }

        json.NewEncoder(w).Encode(record)
    }
}
