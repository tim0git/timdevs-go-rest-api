package handler_update_vehicle_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"timdevs.rest.api.com/m/v2/database"
	"timdevs.rest.api.com/m/v2/handler_update_vehicle"
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
var vinThatDoesNotExist = "GB99999999"
var vinThatDoesExist = "GB000000000"

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
	router.PATCH("/vehicle/:vin", handler_update_vehicle.UpdateVehicle)
	return router
}
func TestReturnsStatusCode200WhenSuccessful(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	request, err := json.Marshal(&mockUpdateVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))

	fmt.Println(requestError)
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestReturnsStatusCode404WhenVehicleVinIsNotPresentInRequest(t *testing.T) {
	t.Parallel()
	router := setupRouter()

	request, err := json.Marshal(&mockUpdateVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", "/vehicle/", bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestReturnsStatusCode400WhenVehicleVinIsNotPresentInDatabase(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	request, err := json.Marshal(&mockUpdateVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesNotExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestReturnsStatusCode400WhenVehicleColorIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Color = ""

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestReturnsStatusCode400WhenVehicleCapacityIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Capacity = vehicle.Capacity{}

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestReturnsStatusCode400WhenVehicleCapacityValueIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Capacity.Value = 0

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestReturnsStatusCode400WhenVehicleCapacityUnitIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Capacity.Unit = ""

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestReturnsStatusCode400WhenVehicleManufacturerIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Manufacturer = ""

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestReturnsStatusCode400WhenVehicleModelIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Model = ""

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestReturnsStatusCode400WhenVehicleModelYearIsNotDefined(t *testing.T) {
	t.Parallel()

	router := setupRouter()

	badMockVehicle := mockUpdateVehicle
	badMockVehicle.Year = 0

	request, err := json.Marshal(&badMockVehicle)
	assert.NoError(t, err)

	req, requestError := http.NewRequest("PATCH", fmt.Sprintf("/vehicle/%s", vinThatDoesExist), bytes.NewBuffer(request))
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
