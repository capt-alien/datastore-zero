package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/capt-alien/datastore-zero/internal/handlers"
	"github.com/stretchr/testify/assert"
)

func TestHireHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hire", nil)
	resp := httptest.NewRecorder()

	handlers.HireHandler(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Expected status 200 OK")
	assert.Contains(t, resp.Body.String(), "Imagine what I’ll do", "Expected motivational message")
	assert.Contains(t, resp.Body.String(), "linkedin", "Expected LinkedIn field")
	assert.Contains(t, resp.Body.String(), "ericbotcher@gmail.com", "Expected email address")
}
