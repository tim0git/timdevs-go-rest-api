package modal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"timdevs.rest.api.com/m/v2/database"
	"timdevs.rest.api.com/m/v2/utils"
)

func PutVehicle(item map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	client := database.DynamoDB()

	putRequest := buildPutItemInput(item)

	return client.PutItem(putRequest)
}

func buildPutItemInput(item map[string]*dynamodb.AttributeValue) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		TableName:                 aws.String(utils.GetTableName()),
		Item:                      item,
		ConditionExpression:       getConditionExpression(),
		ExpressionAttributeNames:  getPutExpressionAttributeNames(),
		ExpressionAttributeValues: getPutExpressionAttributeValues(item),
	}
}

func getConditionExpression() *string {
	return aws.String(`attribute_not_exists(vin) OR (#v <> :val AND attribute_exists(vin))`)
}

func getPutExpressionAttributeNames() map[string]*string {
	return map[string]*string{
		"#v": aws.String("vin"),
	}
}

func getPutExpressionAttributeValues(item map[string]*dynamodb.AttributeValue) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		":val": {
			S: item["vin"].S,
		},
	}
}
