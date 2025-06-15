package main

import (
	"log"
	"net/http"

	"github.com/capt-alien/datastore-zero/internal/handlers"
	"github.com/capt-alien/datastore-zero/internal/db"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Init MariaDB connection
	database := db.InitDB()

	r := chi.NewRouter()
	r.Put("/store/{id}", handlers.PutHandler(database))
	r.Get("/store", handlers.ListHandler(database))
	r.Get("/store/{id}", handlers.GetHandler(database))
	r.Delete("/store/{id}", handlers.DeleteHandler(database))
	r.Get("/hire", handlers.HireHandler)

	log.Println("Server Running on :8080")
	http.ListenAndServe(":8080", r)
}
