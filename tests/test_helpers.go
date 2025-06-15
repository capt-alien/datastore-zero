package tests

import (
	"log"
	"net/http"

	dbmodel "github.com/capt-alien/datastore-zero/internal/db"
	"github.com/capt-alien/datastore-zero/internal/handlers"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MockRouter wires up the PUT handler using /store/{id}
func MockRouter(database *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Put("/store/{id}", handlers.PutHandler(database))
	return r
}

// SetupTestDB sets up an in-memory SQLite DB for testing
func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test DB: %v", err)
	}
	if err := db.AutoMigrate(&dbmodel.Record{}); err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}
	return db
}
