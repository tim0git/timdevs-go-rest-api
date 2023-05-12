package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"net/http"
	"timdevs.rest.api.com/m/v2/error"
	"timdevs.rest.api.com/m/v2/modal"
)

func RetrieveVehicle(c *gin.Context) {
	vin := c.Param("vin")
	vehicle := Vehicle{}

	getVehicleResponse, getVehicleError := modal.GetVehicle(vin)
	if getVehicleError != nil {
		error.DynamoDBError(c, getVehicleError)
		return
	}

	unMarshalError := dynamodbattribute.UnmarshalMap(getVehicleResponse.Item, &vehicle)
	if unMarshalError != nil {
		error.DynamoDBError(c, unMarshalError)
		return
	}

	if vehicle.Vin == "" {
		error.NotFoundError(c)
		return
	}

	c.JSON(http.StatusOK, vehicle)
}
