package health_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/health"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
func TestReturns200StatusCode(t *testing.T) {
	t.Parallel()

	r := setUpRouter()
	r.GET("/health", health.Health)

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestReturnsStatusOK(t *testing.T) {
	t.Parallel()

	expected := health.Status{Status: "OK"}
	expectedJSON, _ := json.Marshal(expected)

	r := setUpRouter()
	r.GET("/health", health.Health)

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	actual, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	assert.Equal(t, string(expectedJSON), string(actual))
}
