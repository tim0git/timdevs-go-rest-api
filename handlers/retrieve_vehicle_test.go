package handlers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/handlers"
)

func setupRetrieveVehicleRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/vehicle/:vin", handlers.RetrieveVehicle)
	return router
}
func TestReturns200StatusCodeWhenVehicleIdIsPassed(t *testing.T) {
	t.Parallel()
	router := setupRetrieveVehicleRouter()

	req, requestError := http.NewRequest("GET", "/vehicle/GB000000000", nil)
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestReturns404StatusCodeWhenVehicleIdIsNotPassed(t *testing.T) {
	t.Parallel()
	router := setupRetrieveVehicleRouter()

	req, requestError := http.NewRequest("GET", "/vehicle/", nil)
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestReturns404StatusCodeWhenVehicleIdIsNotValid(t *testing.T) {
	t.Parallel()
	router := setupRetrieveVehicleRouter()

	req, requestError := http.NewRequest("GET", "/vehicle/NotARealVin", nil)
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
