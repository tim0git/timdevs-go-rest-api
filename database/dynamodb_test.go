package database_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"timdevs.rest.api.com/m/v2/database"
)

var mockRegion string = "us-east-1"
var mockEndpoint string = "http://localhost:8000"

func TestReturnsDynamoDbClient(t *testing.T) {
	t.Parallel()
	actual := database.DynamoDB()
	assert.NotNil(t, actual)
}
func TestReturnsDynamoDbClientWithConfig(t *testing.T) {
	t.Parallel()
	actual := database.DynamoDB()
	assert.NotNil(t, actual.Config)
}
func TestReturnsDynamoDbClientWithRegion(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("AWS_REGION", mockRegion)
	actual := database.DynamoDB()
	assert.Equal(t, mockRegion, *actual.Config.Region)
}
func TestReturnsDynamoDbClientWithEndpoint(t *testing.T) {
	t.Parallel()
	err := os.Setenv("DYNAMODB_ENDPOINT", mockEndpoint)

	actual := database.DynamoDB()

	assert.NoError(t, err)
	assert.Equal(t, mockEndpoint, *actual.Config.Endpoint)
}
