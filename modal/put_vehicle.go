package modal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"timdevs.rest.api.com/m/v2/database"
)

func PutVehicle(item map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	client := database.DynamoDB()

	res, err := client.PutItem(&dynamodb.PutItemInput{
		TableName:           aws.String(os.Getenv("TABLE_NAME")),
		Item:                item,
		ConditionExpression: aws.String(`attribute_not_exists(vin) OR (#v <> :val AND attribute_exists(vin))`),
		ExpressionAttributeNames: map[string]*string{
			"#v": aws.String("vin"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":val": {
				S: aws.String(*item["vin"].S),
			},
		},
	})

	if err != nil {
		return res, err
	}

	return res, nil
}
