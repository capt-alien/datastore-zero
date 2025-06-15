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

func mockRouterDelete(database *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Delete("/store/{id}", handlers.DeleteHandler(database))
	return r
}

func TestDeleteHandler(t *testing.T) {
	db := SetupTestDB()

	// Preload a record to delete
	testRecord := dbmodel.Record{ID: "testy", Value: "mctestface"}
	if err := db.Create(&testRecord).Error; err != nil {
		t.Fatalf("failed to seed test data: %v", err)
	}

	router := mockRouterDelete(db)
	req := httptest.NewRequest(http.MethodDelete, "/store/testy", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status 200 OK")
	assert.Contains(t, resp.Body.String(), "deleted", "Expected confirmation message")
}

func TestDeleteHandler_IDNotFound(t *testing.T) {
	db := SetupTestDB()
	router := mockRouterDelete(db)

	req := httptest.NewRequest(http.MethodDelete, "/store/nonexistentid", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code, "Expected status 404 Not Found")
	assert.Contains(t, resp.Body.String(), "id not found", "Expected error message in response")
}
