package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Put("/store/{key}", putHandler)
	r.Get("/store/{key}", getHandler)
	r.Delete("/store/{key}", deleteHandler)

	log.Println("Server Running on :8080")
	http.ListenAndServe(":8080", r)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PUT /store - not implemented yet"))
}
func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET /store - not implemented yet"))
}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE /store - not implemented yet"))
}
