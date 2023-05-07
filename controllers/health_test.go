package controllers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/controllers"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetHealthReturns200StatusCode(t *testing.T) {
	t.Parallel()
	r := SetUpRouter()
	r.GET("/health", controllers.Health)
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetHealthReturnsStatusOK(t *testing.T) {
	t.Parallel()
	expected, _ := json.Marshal(controllers.HealthStatus{Status: "OK"})
	r := SetUpRouter()
	r.GET("/health", controllers.Health)
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	actual, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, string(expected), string(actual))
}
