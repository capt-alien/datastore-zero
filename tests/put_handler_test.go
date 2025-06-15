package tests


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"github.com/capt-alien/datastore-zero/internal/handlers"

)

func MockRouterPut(database *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Put("/store/{key}", handlers.PutHandler(database))
	return r
}

func TestPutHandler(t *testing.T) {
	db := SetupTestDB()
	router := MockRouter(db)

	reqBody := []byte("test value")
	req := httptest.NewRequest(http.MethodPut, "/store/testkey", bytes.NewReader(reqBody))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Expected status 201 Created")
	assert.Contains(t, resp.Body.String(), "OK", "Expected response to contain 'OK'")
}

func TestPutHandler_BadRequest(t *testing.T) {
	db := SetupTestDB()
	router := MockRouter(db)

	req := httptest.NewRequest(http.MethodPut, "/store/badkey", nil) // nil body == bad request
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code, "Expected status 400 Bad Request")
	assert.Contains(t, resp.Body.String(), "could not read request body", "Expected error message in response")
}
