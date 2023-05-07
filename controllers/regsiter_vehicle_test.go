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

var mockVehicle = controllers.Vehicle{
	Vin:          1234567890,
	Manufacturer: "Tesla",
	Model:        "Model 3",
	Year:         2020,
	Color:        "Red",
	CapacityKwh:  75,
	LicensePlate: "ABC123",
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/vehicle", controllers.RegisterVehicle)
	return router
}
func TestRegisterVehicleReturns201StatusCode(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	request, err := json.Marshal(&mockVehicle)
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

	mockVehicleMissingVin := controllers.Vehicle{
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		CapacityKwh:  mockVehicle.CapacityKwh,
		LicensePlate: mockVehicle.LicensePlate,
	}

	request, err := json.Marshal(mockVehicleMissingVin)
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

	mockVehicleMissingManufacturer := controllers.Vehicle{
		Vin:          mockVehicle.Vin,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		CapacityKwh:  mockVehicle.CapacityKwh,
		LicensePlate: mockVehicle.LicensePlate,
	}
	request, err := json.Marshal(mockVehicleMissingManufacturer)
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

	mockVehicleMissingModel := controllers.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		CapacityKwh:  mockVehicle.CapacityKwh,
		LicensePlate: mockVehicle.LicensePlate,
	}
	request, err := json.Marshal(mockVehicleMissingModel)
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

	mockVehicleMissingYear := controllers.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Color:        mockVehicle.Color,
		CapacityKwh:  mockVehicle.CapacityKwh,
		LicensePlate: mockVehicle.LicensePlate,
	}
	request, err := json.Marshal(mockVehicleMissingYear)
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

	mockVehicleMissingColor := controllers.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		CapacityKwh:  mockVehicle.CapacityKwh,
		LicensePlate: mockVehicle.LicensePlate,
	}
	request, err := json.Marshal(mockVehicleMissingColor)
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

	mockVehicleMissingCapacityKwh := controllers.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		LicensePlate: mockVehicle.LicensePlate,
	}
	request, err := json.Marshal(mockVehicleMissingCapacityKwh)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestRegisterVehicleReturns201StatusCodeWhenLicensePlateIsMissing(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	mockVehicleMissingLicensePlate := controllers.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		CapacityKwh:  mockVehicle.CapacityKwh,
	}
	request, err := json.Marshal(mockVehicleMissingLicensePlate)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
