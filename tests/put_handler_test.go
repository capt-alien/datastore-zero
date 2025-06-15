package tests

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	dbmodel "github.com/capt-alien/datastore-zero/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// mockRouter wires up only the putHandler for isolated testing
func mockRouter(database *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Put("/store/{key}", putHandler(database))
	return r
}

// setupTestDB sets up an in-memory SQLite DB for testing
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to test DB: %v", err)
	}
	if err := db.AutoMigrate(&dbmodel.Record{}); err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}
	return db
}

func TestPutHandler(t *testing.T) {
	db := setupTestDB()
	router := mockRouter(db)

	reqBody := []byte("test value")
	req := httptest.NewRequest(http.MethodPut, "/store/testkey", bytes.NewReader(reqBody))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Expected status 201 Created")
	assert.Contains(t, resp.Body.String(), "OK", "Expected response to contain 'OK'")
}
