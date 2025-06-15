package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	dbmodel "github.com/capt-alien/datastore-zero/internal/db"
	"github.com/capt-alien/datastore-zero/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func mockRouterList(database *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/store", handlers.ListHandler(database))
	return r
}

func TestListHandler(t *testing.T) {
	db := SetupTestDB()

	testRecords := []dbmodel.Record{
		{ID: "testy", Value: "mcTestFace"},
		{ID: "testareno", Value: "faceTest"},
	}
	if err := db.Create(&testRecords).Error; err != nil {
		t.Fatalf("failed to seed test data: %v", err)
	}

	router := mockRouterList(db)
	req := httptest.NewRequest(http.MethodGet, "/store", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status 200 OK")

	var results []dbmodel.Record
	err := json.Unmarshal(resp.Body.Bytes(), &results)
	assert.NoError(t, err, "Response should be valid JSON")
	assert.Len(t, results, 2, "Expected two records in response")
}

func TestListHandler_EmptyDB(t *testing.T) {
	db := SetupTestDB()
	router := mockRouterList(db)

	req := httptest.NewRequest(http.MethodGet, "/store", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status 200 OK")

	var results []dbmodel.Record
	err := json.Unmarshal(resp.Body.Bytes(), &results)
	assert.NoError(t, err, "Response should be valid JSON")
	assert.Len(t, results, 0, "Expected zero records in response")
}
