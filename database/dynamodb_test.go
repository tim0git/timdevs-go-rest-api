package database_test

import (
	"eve.vehicle.api.com/m/v2/database"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var mockRegion string = "us-east-1"
var mockEndpoint string = "http://localhost:8000"

func TestReturnsDynamoDBClient(t *testing.T) {
	t.Parallel()
	actual := database.DynamoDB()
	assert.NotNil(t, actual)
}
func TestReturnsDynamoDBClientWithConfig(t *testing.T) {
	t.Parallel()
	actual := database.DynamoDB()
	assert.NotNil(t, actual.Config)
}
func TestReturnsDynamoDBClientWithRegion(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("AWS_REGION", mockRegion)
	actual := database.DynamoDB()
	assert.Equal(t, mockRegion, *actual.Config.Region)
}
func TestReturnsDynamoDBClientWithEndpoint(t *testing.T) {
	t.Parallel()
	err := os.Setenv("DYNAMODB_ENDPOINT", mockEndpoint)

	actual := database.DynamoDB()

	assert.NoError(t, err)
	assert.Equal(t, mockEndpoint, *actual.Config.Endpoint)
}
