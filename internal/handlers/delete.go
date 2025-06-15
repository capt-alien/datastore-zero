package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/capt-alien/datastore-zero/internal/db"
    "github.com/go-chi/chi/v5"
    "gorm.io/gorm"
)

func DeleteHandler(database *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        id := chi.URLParam(r, "id")
        if id == "" {
            JSONError(w, "missing id", http.StatusBadRequest)
            return
        }

        var record db.Record
        err := database.First(&record, "id = ?", id).Error
        if err != nil {
            JSONError(w, "record not found", http.StatusNotFound)
            return
        }

        if err := database.Delete(&record).Error; err != nil {
            JSONError(w, "failed to delete record", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
    }
}
