package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"

	"timdevs.rest.api.com/m/v2/controllers"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/vehicle", controllers.RegisterVehicle)
	return router
}
func TestRegisterVehicleReturns201StatusCode(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	mockVehicle := controllers.Vehicle{
		Vin:          1234567890,
		Manufacturer: "Tesla",
		Model:        "Model 3",
		Year:         2020,
		Color:        "Red",
		CapacityKwh:  75,
		LicensePlate: "ABC123",
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestRegisterVehicleReturnsValidationErrorWhenVinIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Vin' Error:Field validation for 'Vin' failed on the 'required' tag",
	}

	expected, err := json.Marshal(validationError)
	assert.NoError(t, err)

	mockVehicle := controllers.Vehicle{
		Manufacturer: "Tesla",
		Model:        "Model 3",
		Year:         2020,
		Color:        "Red",
		CapacityKwh:  75,
		LicensePlate: "ABC123",
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestRegisterVehicleReturnsValidationErrorWhenManufacturerIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Manufacturer' Error:Field validation for 'Manufacturer' failed on the 'required' tag",
	}

	expected, err := json.Marshal(validationError)
	assert.NoError(t, err)

	mockVehicle := controllers.Vehicle{
		Vin:          1234567890,
		Model:        "Model 3",
		Year:         2020,
		Color:        "Red",
		CapacityKwh:  75,
		LicensePlate: "ABC123",
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestRegisterVehicleReturnsValidationErrorWhenModelIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Model' Error:Field validation for 'Model' failed on the 'required' tag",
	}

	expected, err := json.Marshal(validationError)
	assert.NoError(t, err)

	mockVehicle := controllers.Vehicle{
		Vin:          1234567890,
		Manufacturer: "Tesla",
		Year:         2020,
		Color:        "Red",
		CapacityKwh:  75,
		LicensePlate: "ABC123",
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestRegisterVehicleReturnsValidationErrorWhenYearIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Year' Error:Field validation for 'Year' failed on the 'required' tag",
	}

	expected, err := json.Marshal(validationError)
	assert.NoError(t, err)

	mockVehicle := controllers.Vehicle{
		Vin:          1234567890,
		Manufacturer: "Tesla",
		Model:        "Model 3",
		Color:        "Red",
		CapacityKwh:  75,
		LicensePlate: "ABC123",
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestRegisterVehicleReturnsValidationErrorWhenColorIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Color' Error:Field validation for 'Color' failed on the 'required' tag",
	}

	expected, err := json.Marshal(validationError)
	assert.NoError(t, err)

	mockVehicle := controllers.Vehicle{
		Vin:          1234567890,
		Manufacturer: "Tesla",
		Model:        "Model 3",
		Year:         2020,
		CapacityKwh:  75,
		LicensePlate: "ABC123",
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestRegisterVehicleReturnsValidationErrorWhenCapacityKwhIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.CapacityKwh' Error:Field validation for 'CapacityKwh' failed on the 'required' tag",
	}

	expected, err := json.Marshal(validationError)
	assert.NoError(t, err)

	mockVehicle := controllers.Vehicle{
		Vin:          1234567890,
		Manufacturer: "Tesla",
		Model:        "Model 3",
		Year:         2020,
		Color:        "Red",
		LicensePlate: "ABC123",
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestRegisterVehicleWhenLicensePlateIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	mockVehicle := controllers.Vehicle{
		Vin:          1234567890,
		Manufacturer: "Tesla",
		Model:        "Model 3",
		Year:         2020,
		Color:        "Red",
		CapacityKwh:  75,
	}
	request, err := json.Marshal(mockVehicle)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
