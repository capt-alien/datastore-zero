package tests


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


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
