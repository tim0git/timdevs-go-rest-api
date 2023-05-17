package handlers

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
	"net/http"
	"timdevs.rest.api.com/m/v2/error"
	"timdevs.rest.api.com/m/v2/modal"
)

type UpdateVehicleRequest struct {
	Vin          string          `json:"vin"`
	Manufacturer string          `json:"manufacturer"`
	Model        string          `json:"model"`
	Year         int             `json:"year"`
	Color        string          `json:"color"`
	Capacity     VehicleCapacity `json:"capacity"`
	LicensePlate string          `json:"license_plate"`
}

func UpdateVehicle(c *gin.Context) {
	vehicle := UpdateVehicleRequest{}

	validationError := c.ShouldBindJSON(&vehicle)
	if validationError != nil {
		error.ValidationError(c, validationError)
		return
	}

	marshalledVehicle, marshalError := dynamodbattribute.MarshalMap(&vehicle)
	if marshalError != nil {
		error.DynamoDBError(c, marshalError)
		return
	}

	_, updateItemError := modal.UpdateVehicle(marshalledVehicle)
	if updateItemError != nil {
		error.DynamoDBError(c, updateItemError)
		return
	}

	c.Status(http.StatusOK)
}
