package controllers_test

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"timdevs.rest.api.com/m/v2/controllers"
	dynamodb2 "timdevs.rest.api.com/m/v2/database"
	"timdevs.rest.api.com/m/v2/validators"
)

var mockVehicle = validators.Vehicle{
	Vin:          "GB29HP0K456785",
	Manufacturer: "Tesla",
	Model:        "Model 3",
	Year:         2020,
	Color:        "Red",
	Capacity: validators.VehicleCapacity{
		Value: 75,
		Unit:  "kWh",
	},
	LicensePlate: "ABC123",
}

func TestMain(m *testing.M) {
	tableName := fmt.Sprintf("Vehicles-%v", random.UniqueId())

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "mock-key")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "mock-secret")
	_ = os.Setenv("DYNAMODB_ENDPOINT", "http://localhost:8000")
	_ = os.Setenv("TABLE_NAME", tableName)

	client := dynamodb2.Client()
	_, err := client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("vin"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("vin"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	assert.NoError(m, err)

	// Run the tests.
	exitCode := m.Run()

	// Delete the table after running all the tests.
	_, err = client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	assert.NoError(m, err)

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "")
	_ = os.Setenv("DYNAMODB_ENDPOINT", "")
	_ = os.Setenv("TABLE_NAME", "")

	os.Exit(exitCode)
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

	mockVehicleMissingVin := validators.Vehicle{
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		Capacity:     mockVehicle.Capacity,
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

	mockVehicleMissingManufacturer := validators.Vehicle{
		Vin:          mockVehicle.Vin,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		Capacity:     mockVehicle.Capacity,
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

	mockVehicleMissingModel := validators.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		Capacity:     mockVehicle.Capacity,
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

	mockVehicleMissingYear := validators.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Color:        mockVehicle.Color,
		Capacity:     mockVehicle.Capacity,
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

	mockVehicleMissingColor := validators.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Capacity:     mockVehicle.Capacity,
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
func TestRegisterVehicleReturnsValidationErrorWhenCapacityIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Capacity.Value' Error:Field validation for 'Value' failed on the 'required' tag\nKey: 'Vehicle.Capacity.Unit' Error:Field validation for 'Unit' failed on the 'required' tag",
	}

	expected, err := json.Marshal(validationError)
	assert.NoError(t, err)

	mockVehicleMissingCapacityKwh := validators.Vehicle{
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

	mockVehicleMissingLicensePlate := validators.Vehicle{
		Vin:          mockVehicle.Vin,
		Manufacturer: mockVehicle.Manufacturer,
		Model:        mockVehicle.Model,
		Year:         mockVehicle.Year,
		Color:        mockVehicle.Color,
		Capacity:     mockVehicle.Capacity,
	}
	request, err := json.Marshal(mockVehicleMissingLicensePlate)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
