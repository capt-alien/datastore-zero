package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/capt-alien/datastore-zero/internal/db"
    "gorm.io/gorm"
)


func ListHandler(database *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var records []db.Record
		err := database.Find(&records).Error
		if err != nil {
			http.Error(w, "failed to fetch records", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)

	}
}
