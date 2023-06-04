package modal

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"timdevs.rest.api.com/m/v2/database"
	"timdevs.rest.api.com/m/v2/vehicle"
)

func UpdateVehicle(vehicle vehicle.Update, vin string) (*dynamodb.UpdateItemOutput, error) {
	client := database.DynamoDB()

	updateRequest := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]*dynamodb.AttributeValue{
			"vin": {
				S: aws.String(vin),
			},
		},
	}

	updateRequest.ConditionExpression = aws.String(`attribute_exists(vin)`)

	updateRequest.UpdateExpression = aws.String("set #manufacturer = :manufacturer, #model = :model, #year = :year, #color = :color, #batteryCapacity = :batteryCapacity")

	updateRequest.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{
		":manufacturer": {
			S: aws.String(vehicle.Manufacturer),
		},
		":model": {
			S: aws.String(vehicle.Model),
		},
		":year": {
			N: aws.String(fmt.Sprintf("%d", vehicle.Year)),
		},
		":color": {
			S: aws.String(vehicle.Color),
		},
		":batteryCapacity": {
			M: map[string]*dynamodb.AttributeValue{
				"value": {
					N: aws.String(fmt.Sprintf("%d", vehicle.Capacity.Value)),
				},
				"unit": {
					S: aws.String(fmt.Sprintf("%s", vehicle.Capacity.Unit)),
				},
			},
		},
	}

	updateRequest.ExpressionAttributeNames = map[string]*string{
		"#model":           aws.String("model"),
		"#year":            aws.String("year"),
		"#color":           aws.String("color"),
		"#batteryCapacity": aws.String("capacity"),
		"#manufacturer":    aws.String("manufacturer"),
	}

	return client.UpdateItem(updateRequest)
}
