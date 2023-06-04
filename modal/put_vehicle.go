package modal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"timdevs.rest.api.com/m/v2/database"
)

func PutVehicle(item map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	client := database.DynamoDB()

	putRequest := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item:      item,
	}

	putRequest.ConditionExpression = aws.String(`attribute_not_exists(vin) OR (#v <> :val AND attribute_exists(vin))`)

	putRequest.ExpressionAttributeNames = map[string]*string{
		"#v": aws.String("vin"),
	}
	putRequest.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
		":val": {
			S: aws.String(*item["vin"].S),
		},
	}

	return client.PutItem(putRequest)
}
