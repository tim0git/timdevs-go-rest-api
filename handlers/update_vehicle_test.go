package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/handlers"
	"timdevs.rest.api.com/m/v2/vehicle"
)

var mockUpdateVehicle = vehicle.Update{
	Manufacturer: "Tesla",
	Model:        "Model 3",
	Year:         2020,
	Color:        "Red",
	Capacity: vehicle.Capacity{
		Value: 75,
		Unit:  "kWh",
	},
	LicensePlate: "ABC123",
}

func setupUpdateVehicleRouter() *gin.Engine {
	router := gin.Default()
	router.PATCH("/vehicle/:vin", handlers.UpdateVehicle)
	return router
}

func TestUpdateVehicleReturnsStatusCode200(t *testing.T) {
	t.Parallel()
	router := setupUpdateVehicleRouter()

	request, err := json.Marshal(&mockUpdateVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", "/vehicle/GB000000000", bytes.NewBuffer(request))

	fmt.Println(requestError)
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateVehicleReturnsStatusCode404WhenVehicleVinIsNotPresentInRequest(t *testing.T) {
	t.Parallel()
	router := setupUpdateVehicleRouter()

	request, err := json.Marshal(&mockUpdateVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", "/vehicle/", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateVehicleReturnsStatusCode400WhenVehicleColorIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupUpdateVehicleRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Color = ""

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", "/vehicle/GB000000000", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
