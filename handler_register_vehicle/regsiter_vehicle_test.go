package handler_register_vehicle_test

import (
	"bytes"
	"eve.vehicle.api.com/m/v2/database"
	"eve.vehicle.api.com/m/v2/handler_register_vehicle"
	"eve.vehicle.api.com/m/v2/vehicle"
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
)

var vinThatDoesExist = "GB000000000"
var vinThatDoesNotExist = "GB99999999"
var mockVehicle = vehicle.Vehicle{
	Vin:          vinThatDoesNotExist,
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

func TestMain(m *testing.M) {
	tableName := fmt.Sprintf("Vehicles-%v", random.UniqueId())

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "mock-key")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "mock-secret")
	_ = os.Setenv("DYNAMODB_ENDPOINT", "http://localhost:8000")
	_ = os.Setenv("TABLE_NAME", tableName)

	// Create the table before running the tests.
	client := database.DynamoDB()
	_, createTableError := client.CreateTable(&dynamodb.CreateTableInput{
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
	if createTableError != nil {
		panic(createTableError)
	}

	// Add an item to the table.
	_, putItemError := client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"vin": {
				S: aws.String(vinThatDoesExist),
			},
		},
	})
	if putItemError != nil {
		panic(putItemError)
	}

	// Run the tests.
	exitCode := m.Run()

	// Delete the table after running all the tests.
	_, deleteTableError := client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	if deleteTableError != nil {
		panic(deleteTableError)
	}

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "")
	_ = os.Setenv("DYNAMODB_ENDPOINT", "")
	_ = os.Setenv("TABLE_NAME", "")

	os.Exit(exitCode)
}
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/vehicle", handler_register_vehicle.RegisterVehicle)
	return router
}
func TestReturns201StatusCodeWhenAllFieldsArePresent(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	request, err := json.Marshal(&mockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestReturns201StatusCodeWhenLicensePlateIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	mockVehicleMissingLicensePlate := mockVehicle
	mockVehicleMissingLicensePlate.Vin = fmt.Sprintf(random.UniqueId())
	mockVehicleMissingLicensePlate.LicensePlate = ""
	request, marshalError := json.Marshal(&mockVehicleMissingLicensePlate)
	assert.NoError(t, marshalError)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestReturnsValidationErrorWhenVinIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Vin' Error:Field validation for 'Vin' failed on the 'required' tag",
	}

	expected, err := json.Marshal(&validationError)
	assert.NoError(t, err)

	mockVehicleMissingVin := mockVehicle
	mockVehicleMissingVin.Vin = ""

	request, marshalError := json.Marshal(mockVehicleMissingVin)
	assert.NoError(t, marshalError)
	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestReturnsValidationErrorWhenManufacturerIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Manufacturer' Error:Field validation for 'Manufacturer' failed on the 'required' tag",
	}

	expected, err := json.Marshal(&validationError)
	assert.NoError(t, err)

	mockVehicleMissingManufacturer := mockVehicle
	mockVehicleMissingManufacturer.Manufacturer = ""
	request, marshalError := json.Marshal(mockVehicleMissingManufacturer)
	assert.NoError(t, marshalError)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestReturnsValidationErrorWhenModelIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Model' Error:Field validation for 'Model' failed on the 'required' tag",
	}

	expected, err := json.Marshal(&validationError)
	assert.NoError(t, err)

	mockVehicleMissingModel := mockVehicle
	mockVehicleMissingModel.Model = ""
	request, marshalError := json.Marshal(mockVehicleMissingModel)
	assert.NoError(t, marshalError)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestReturnsValidationErrorWhenYearIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Year' Error:Field validation for 'Year' failed on the 'required' tag",
	}

	expected, err := json.Marshal(&validationError)
	assert.NoError(t, err)

	mockVehicleMissingYear := mockVehicle
	mockVehicleMissingYear.Year = 0
	request, marshalError := json.Marshal(mockVehicleMissingYear)
	assert.NoError(t, marshalError)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestReturnsValidationErrorWhenColorIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Color' Error:Field validation for 'Color' failed on the 'required' tag",
	}

	expected, err := json.Marshal(&validationError)
	assert.NoError(t, err)

	mockVehicleMissingColor := mockVehicle
	mockVehicleMissingColor.Color = ""
	request, marshalError := json.Marshal(mockVehicleMissingColor)
	assert.NoError(t, marshalError)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestReturnsValidationErrorWhenCapacityIsMissing(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	validationError := gin.H{
		"error":   "VALIDATEERR-1",
		"message": "Key: 'Vehicle.Capacity.Value' Error:Field validation for 'Value' failed on the 'required' tag\nKey: 'Vehicle.Capacity.Unit' Error:Field validation for 'Unit' failed on the 'required' tag",
	}

	expected, err := json.Marshal(&validationError)
	assert.NoError(t, err)

	mockVehicleMissingCapacityKwh := mockVehicle
	mockVehicleMissingCapacityKwh.Capacity = vehicle.Capacity{}
	request, marshalError := json.Marshal(&mockVehicleMissingCapacityKwh)
	assert.NoError(t, marshalError)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
func TestReturnsStatusCode400WhenAVehicleAlreadyExists(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	mockVehicleAlreadyExists := mockVehicle
	mockVehicleAlreadyExists.Vin = vinThatDoesExist
	request, marshalError := json.Marshal(&mockVehicleAlreadyExists)
	assert.NoError(t, marshalError)

	dynamoError := gin.H{
		"error":   "DYNAMOERR-1",
		"message": "ConditionalCheckFailedException",
	}
	expected, err := json.Marshal(&dynamoError)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("POST", "/vehicle", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
