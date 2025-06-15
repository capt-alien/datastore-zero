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

        key := chi.URLParam(r, "key")
        if key == "" {
            JSONError(w, "missing key", http.StatusBadRequest)
            return
        }

        var record db.Record
        err := database.First(&record, "key = ?", key).Error
        if err != nil {
            JSONError(w, "key not found", http.StatusNotFound)
            return
        }

        if err := database.Delete(&record).Error; err != nil {
            JSONError(w, "failed to delete record", http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
    }
}
