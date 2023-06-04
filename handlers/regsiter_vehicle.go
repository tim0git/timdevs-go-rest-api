package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"net/http"
	"timdevs.rest.api.com/m/v2/error"
	"timdevs.rest.api.com/m/v2/modal"
	"timdevs.rest.api.com/m/v2/vehicle"
)

func RegisterVehicle(c *gin.Context) {
	newVehicle := vehicle.Vehicle{}

	validationError := c.ShouldBindJSON(&newVehicle)
	if validationError != nil {
		error.ValidationError(c, validationError)
		return
	}

	marshalledVehicle, marshalError := dynamodbattribute.MarshalMap(&newVehicle)
	if marshalError != nil {
		error.DynamoDBError(c, marshalError)
		return
	}

	_, putItemError := modal.PutVehicle(marshalledVehicle)
	if putItemError != nil {
		error.DynamoDBError(c, putItemError)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
