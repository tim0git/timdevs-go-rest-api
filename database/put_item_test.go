package database_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	dynamodb2 "timdevs.rest.api.com/m/v2/database"
)

func TestMain(m *testing.M) {
	tableName := fmt.Sprintf("Test-%v", random.UniqueId())

	_ = os.Setenv("AWS_ACCESS_KEY_ID", "mock-key")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "mock-secret")
	_ = os.Setenv("DYNAMODB_ENDPOINT", "http://localhost:8000")
	_ = os.Setenv("TABLE_NAME", tableName)

	client := dynamodb2.Client()
	_, err := client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(tableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
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

func TestPutsItemWithoutError(t *testing.T) {
	item := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String("123"),
		},
	}

	_, err := dynamodb2.PutItem(item)
	assert.NoError(t, err)
}
