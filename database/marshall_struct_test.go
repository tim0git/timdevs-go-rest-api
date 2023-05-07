package database_test

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"timdevs.rest.api.com/m/v2/validators"
)

func TestMarshallStructReturnsCorrectlyMarshalledItem(t *testing.T) {
	body := validators.Vehicle{
		Vin:          "GB29HP0K456785",
		Manufacturer: "Tesla",
		Model:        "Model 3",
		Year:         2020,
		Color:        "Red",
		Capacity: validators.VehicleCapacity{
			Value: 75,
			Unit:  "kWh",
		},
		LicensePlate: "ABC123",
	}

	item, err := dynamodbattribute.MarshalMap(body)
	assert.NoError(t, err)

	expectedYearString := strconv.Itoa(body.Year)
	expectedCapacityValueString := strconv.Itoa(body.Capacity.Value)

	expect := map[string]*dynamodb.AttributeValue{
		"vin": {
			S: &body.Vin,
		},
		"manufacturer": {
			S: &body.Manufacturer,
		},
		"model": {
			S: &body.Model,
		},
		"year": {
			N: &expectedYearString,
		},
		"color": {
			S: &body.Color,
		},
		"capacity": {
			M: map[string]*dynamodb.AttributeValue{
				"value": {
					N: &expectedCapacityValueString,
				},
				"unit": {
					S: &body.Capacity.Unit,
				},
			},
		},
		"license_plate": {
			S: &body.LicensePlate,
		},
	}

	assert.Equal(t, expect, item)
}
