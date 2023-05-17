package modal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"timdevs.rest.api.com/m/v2/database"
)

func UpdateVehicle(item map[string]*dynamodb.AttributeValue) (*dynamodb.UpdateItemOutput, error) {
	client := database.DynamoDB()

	input := &dynamodb.UpdateItemInput{
		Key:          item,
		ReturnValues: aws.String("UPDATED_NEW"),
		TableName:    aws.String(os.Getenv("TABLE_NAME")),
	}

	return client.UpdateItem(input)
}
