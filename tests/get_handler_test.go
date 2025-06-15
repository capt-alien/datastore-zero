package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	dbmodel "github.com/capt-alien/datastore-zero/internal/db"
	"github.com/capt-alien/datastore-zero/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// mockRouterGet wires up only the GetHandler for isolated testing
func MockRouterGet(database *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/store/{key}", handlers.GetHandler(database))
	return r
}

func TestGetHandler(t *testing.T) {
	db := SetupTestDB()

	// Preload test record
	testRecord := dbmodel.Record{Key: "testkey", Value: "testvalue"}
	if err := db.Create(&testRecord).Error; err != nil {
		t.Fatalf("failed to seed test data: %v", err)
	}

	router := MockRouterGet(db)

	req := httptest.NewRequest(http.MethodGet, "/store/testkey", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status 200 OK")
	assert.Contains(t, resp.Body.String(), "testvalue", "Expected response to contain 'testvalue'")
}
