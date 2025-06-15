package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/capt-alien/datastore-zero/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/go-chi/chi/v5"
)

func TestFullLifecycle(t *testing.T) {
	db := SetupTestDB()
	r := chi.NewRouter()

	r.Put("/store/{id}", handlers.PutHandler(db))
	r.Get("/store/{id}", handlers.GetHandler(db))
	r.Delete("/store/{id}", handlers.DeleteHandler(db))

	// Step 1: PUT
	putReq := httptest.NewRequest(http.MethodPut, "/store/e2eid", bytes.NewBufferString("e2evalue"))
	putResp := httptest.NewRecorder()
	r.ServeHTTP(putResp, putReq)
	assert.Equal(t, http.StatusCreated, putResp.Code, "PUT should return 201 Created")

	// Step 2: GET
	getReq := httptest.NewRequest(http.MethodGet, "/store/e2eid", nil)
	getResp := httptest.NewRecorder()
	r.ServeHTTP(getResp, getReq)
	assert.Equal(t, http.StatusOK, getResp.Code, "GET should return 200 OK")
	assert.Contains(t, getResp.Body.String(), "e2evalue", "GET response should contain correct value")

	// Step 3: DELETE
	deleteReq := httptest.NewRequest(http.MethodDelete, "/store/e2eid", nil)
	deleteResp := httptest.NewRecorder()
	r.ServeHTTP(deleteResp, deleteReq)
	assert.Equal(t, http.StatusOK, deleteResp.Code, "DELETE should return 200 OK")
	assert.Contains(t, deleteResp.Body.String(), "deleted", "DELETE response should confirm deletion")

	// Step 4: GET again (should be gone)
	getReqAgain := httptest.NewRequest(http.MethodGet, "/store/e2eid", nil)
	getRespAgain := httptest.NewRecorder()
	r.ServeHTTP(getRespAgain, getReqAgain)
	assert.Equal(t, http.StatusNotFound, getRespAgain.Code, "GET after DELETE should return 404")
	assert.Contains(t, getRespAgain.Body.String(), "record not found", "Expected error message after deletion")
}
