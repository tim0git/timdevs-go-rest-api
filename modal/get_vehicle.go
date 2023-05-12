package modal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"timdevs.rest.api.com/m/v2/database"
)

func GetVehicle(vin string) (*dynamodb.GetItemOutput, error) {
	client := database.DynamoDB()

	input := &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"vin": {
				S: aws.String(vin),
			},
		},
	}

	return client.GetItem(input)
}
