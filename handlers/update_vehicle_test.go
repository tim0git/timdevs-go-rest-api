package handlers_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"timdevs.rest.api.com/m/v2/handlers"
)

var vehicle = handlers.Vehicle{
	//Vin:          "GB000000000",
	Manufacturer: "Audi",
	Model:        "A4",
	Year:         2018,
	Color:        "Black",
	Capacity: handlers.VehicleCapacity{
		Value: 14,
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

	request, err := json.Marshal(&vehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", "/vehicle/GB000000000", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

//func TestUpdateVehicleReturnsStatusCode404WhenVehicleIdIsNotFound(t *testing.T) {
//	t.Parallel()
//	router := setupUpdateVehicleRouter()
//
//	request, err := json.Marshal(&vehicle)
//	assert.NoError(t, err)
//
//	req, requestError := http.NewRequest("PATCH", "/vehicle/NotARealVin", bytes.NewBuffer(request))
//	assert.NoError(t, requestError)
//
//	w := httptest.NewRecorder()
//	router.ServeHTTP(w, req)
//
//	assert.Equal(t, http.StatusNotFound, w.Code)
//}
