package database

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"timdevs.rest.api.com/m/v2/validators"
)

func MarshallStruct(body validators.Vehicle) (map[string]*dynamodb.AttributeValue, error) {
	item, err := dynamodbattribute.MarshalMap(body)

	if err != nil {
		return nil, err
	}

	return item, nil
}
