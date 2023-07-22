package modal_test

import (
	"eve.vehicle.api.com/m/v2/database"
	"eve.vehicle.api.com/m/v2/modal"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tableName := fmt.Sprintf("Vehicles-%v", random.UniqueId())

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "mock-key")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "mock-secret")
	_ = os.Setenv("DYNAMODB_ENDPOINT", "http://localhost:8000")
	_ = os.Setenv("TABLE_NAME", tableName)

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
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
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
func TestPutsVehicleWithoutError(t *testing.T) {
	t.Parallel()
	item := map[string]*dynamodb.AttributeValue{
		"vin": {
			S: aws.String("123"),
		},
	}

	_, err := modal.PutVehicle(item)
	assert.NoError(t, err)
}
