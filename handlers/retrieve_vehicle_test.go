package handlers_test

import (
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
	"timdevs.rest.api.com/m/v2/handlers"
)

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
				S: aws.String("GB000000000"),
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

func setupRetrieveVehicleRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/vehicle/:vin", handlers.RetrieveVehicle)
	return router
}
func TestReturns200StatusCodeWhenVehicleIsFound(t *testing.T) {
	t.Parallel()
	router := setupRetrieveVehicleRouter()

	req, requestError := http.NewRequest("GET", "/vehicle/GB000000000", nil)
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestReturns404StatusCodeWhenVehicleIdIsNotFound(t *testing.T) {
	t.Parallel()
	router := setupRetrieveVehicleRouter()

	req, requestError := http.NewRequest("GET", "/vehicle/NotARealVin", nil)
	assert.NoError(t, requestError)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
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
