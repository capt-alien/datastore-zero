package main

import (
	"io"
	"log"
	"net/http"

	"encoding/json"
	"github.com/capt-alien/datastore-zero/internal/db"
	"github.com/capt-alien/datastore-zero/internal/handlers"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func main() {
	// Init SQLite
	database := db.InitDB("./data/data.db")

	r := chi.NewRouter()

	r.Put("/store/{key}", putHandler(database))
	r.Get("/store", listHandler(database))
	r.Get("/store/{key}", getHandler(database))
	r.Delete("/store/{key}", deleteHandler(database))
	r.Get("/hire", HireHandler)

	log.Println("Server Running on :8080")
	http.ListenAndServe(":8080", r)
}

func putHandler(database *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key := chi.URLParam(r, "key")
		valueBytes, err := io.ReadAll(r.Body)
		if err != nil {
			handlers.JSONError(w, "could not read request body", http.StatusBadRequest)
			return
		}

		value := string(valueBytes)
		record := db.Record{Key: key, Value: value}

		if err := database.Save(&record).Error; err != nil {
			handlers.JSONError(w, "failed to save record", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	}
}

func getHandler(database *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key := chi.URLParam(r, "key")

		var record db.Record
		err := database.First(&record, "key = ?", key).Error
		if err != nil {
			handlers.JSONError(w, "key not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
	}
}

func deleteHandler(database *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		key := chi.URLParam(r, "key")
		if key == "" {
			handlers.JSONError(w, "missing key", http.StatusBadRequest)
			return
		}

		var record db.Record
		err := database.First(&record, "key = ?", key).Error
		if err != nil {
			handlers.JSONError(w, "key not found", http.StatusNotFound)
			return
		}

		if err := database.Delete(&record).Error; err != nil {
			handlers.JSONError(w, "failed to delete record", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
	}
}

func listHandler(database *gorm.DB) http.HandlerFunc {
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

func HireHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := map[string]string{
		"message": "Hey ASE, this took less than 24 hours to build. Imagine what I’ll do with a badge and a paycheck.",
	}
	json.NewEncoder(w).Encode(message)
}
