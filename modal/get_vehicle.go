package modal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"timdevs.rest.api.com/m/v2/database"
	"timdevs.rest.api.com/m/v2/utils"
)

func GetVehicle(vin string) (*dynamodb.GetItemOutput, error) {
	client := database.DynamoDB()

	input := buildGetItemInput(vin)

	return client.GetItem(input)
}

func buildGetItemInput(vin string) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		TableName: aws.String(utils.GetTableName()),
		Key: map[string]*dynamodb.AttributeValue{
			"vin": {
				S: aws.String(vin),
			},
		},
	}
}
